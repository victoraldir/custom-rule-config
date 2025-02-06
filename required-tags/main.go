package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
)

var (
	region        = "eu-central-1"
	sess          *session.Session
	configService ConfigRepository
	database      DynamoDBRepository
	resourceGroup ResourceGroupRepository
	ctx           = context.Background()
)

func init() {
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: aws.Config{
			Region: aws.String(region),
		},
	}))
	configService = NewConfigRepository(sess)
	database = NewDynamoDBRepository(sess)
	resourceGroup = NewResourceGroupRepository(sess)
}

func handler(request events.ConfigEvent) (string, error) {
	printJSON(request)

	var invokingEvent InvokingEvent
	if err := json.Unmarshal([]byte(request.InvokingEvent), &invokingEvent); err != nil {
		return handleError("Error unmarshalling invoking event", err)
	}

	log.Println("ConfigurationItemDiff:", invokingEvent.ConfigurationItemDiff)

	objectCatalog, err := database.GetItem(invokingEvent.ConfigurationItem.Tags.ObjectId)

	if err != nil || objectCatalog == nil {
		if err := PutEvaluations(ctx, invokingEvent.ConfigurationItem.ResourceID, invokingEvent.ConfigurationItem.ResourceType, "NON_COMPLIANT", "Missing required tags", request.ResultToken, invokingEvent.NotificationCreationTime); err != nil {
			return handleError("Error putting evaluations:", err)
		}
		return "", err
	}

	if !invokingEvent.ConfigurationItem.Tags.IsCompliant() {
		return handleNonCompliant(invokingEvent, request, *objectCatalog)
	}

	if !objectCatalog.IsCompliant(invokingEvent.ConfigurationItem.Tags) {
		return handleNonCompliant(invokingEvent, request, *objectCatalog)
	}

	return handleCompliant(invokingEvent, request)
}

func handleNonCompliant(invokingEvent InvokingEvent, request events.ConfigEvent, objectCatalog CatalogObject) (string, error) {

	if err := Remediate(objectCatalog, invokingEvent.ConfigurationItem.ARN); err != nil {
		log.Println("Error remediating:", err)
		if err := PutEvaluations(ctx, invokingEvent.ConfigurationItem.ResourceID, invokingEvent.ConfigurationItem.ResourceType, "NON_COMPLIANT", "Missing required tags", request.ResultToken, invokingEvent.NotificationCreationTime); err != nil {
			log.Println("Error putting evaluations:", err)
		}
		return "", err
	}

	if err := PutEvaluations(ctx, invokingEvent.ConfigurationItem.ResourceID, invokingEvent.ConfigurationItem.ResourceType, "COMPLIANT", "Resource is compliant", request.ResultToken, invokingEvent.NotificationCreationTime); err != nil {
		return handleError("Error putting evaluations", err)
	}

	return "CONFORMANT", nil
}

func handleCompliant(invokingEvent InvokingEvent, request events.ConfigEvent) (string, error) {
	if err := PutEvaluations(ctx, invokingEvent.ConfigurationItem.ResourceID, invokingEvent.ConfigurationItem.ResourceType, "COMPLIANT", "Resource is compliant", request.ResultToken, invokingEvent.NotificationCreationTime); err != nil {
		return handleError("Error putting evaluations", err)
	}
	return "CONFORMANT", nil
}

func Remediate(objectCatalog CatalogObject, arn string) error {
	_, err := resourceGroup.TagResources(&resourcegroupstaggingapi.TagResourcesInput{
		ResourceARNList: []*string{aws.String(arn)},
		Tags:            objectCatalog.ToTags(),
	})
	return err
}

func PutEvaluations(ctx context.Context, resourceId, resourceType, complianceType, annotation, resultToken string, notificationCreationTime time.Time) error {
	output, err := configService.PutEvaluations(ctx, &configservice.PutEvaluationsInput{
		Evaluations: []*configservice.Evaluation{
			{
				ComplianceResourceId:   aws.String(resourceId),
				ComplianceResourceType: aws.String(resourceType),
				ComplianceType:         aws.String(complianceType),
				Annotation:             aws.String(annotation),
				OrderingTimestamp:      aws.Time(notificationCreationTime),
			},
		},
		ResultToken: aws.String(resultToken),
	})
	if err != nil {
		return err
	}
	log.Println("PutEvaluations output:", output)
	return nil
}

func printJSON(v interface{}) {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Println("Error marshalling to JSON:", err)
		return
	}
	log.Println(string(jsonData))
}

func handleError(message string, err error) (string, error) {
	log.Println(message, err)
	return "", err
}

func main() {
	lambda.Start(handler)
}
