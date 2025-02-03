package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.ConfigEvent) (string, error) {

	fmt.Println("ConfigRuleName: ", request.ConfigRuleName)
	fmt.Println("ConfigRuleArn: ", request.ConfigRuleArn)
	fmt.Println("ConfigRuleId: ", request.ConfigRuleID)
	fmt.Println("AccountId: ", request.AccountID)

	return "CONFORMANT", nil
}

func PutEvaluations(configItem, resultToken, complianceType, remediationStatus string) error {
	fmt.Println("TODO Putting evaluations")
	return nil
}

func Remediate() (string, error) {
	fmt.Println("TODO Remediating noncompliant resources")
	return "REMEDIATED", nil
}

func main() {
	lambda.Start(handler)
}
