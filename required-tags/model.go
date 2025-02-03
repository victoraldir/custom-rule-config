package main

import "time"

type InvokingEvent struct {
	ConfigurationItemDiff    interface{}       `json:"configurationItemDiff"`
	ConfigurationItem        ConfigurationItem `json:"configurationItem"`
	NotificationCreationTime time.Time         `json:"notificationCreationTime"`
	MessageType              string            `json:"messageType"`
	RecordVersion            string            `json:"recordVersion"`
}

type ConfigurationItem struct {
	RelatedEvents                 []interface{}              `json:"relatedEvents"`
	Relationships                 []interface{}              `json:"relationships"`
	Configuration                 Configuration              `json:"configuration"`
	SupplementaryConfiguration    SupplementaryConfiguration `json:"supplementaryConfiguration"`
	Tags                          Tags                       `json:"tags"`
	ConfigurationItemVersion      string                     `json:"configurationItemVersion"`
	ConfigurationItemCaptureTime  time.Time                  `json:"configurationItemCaptureTime"`
	ConfigurationStateID          int64                      `json:"configurationStateId"`
	AwsAccountID                  string                     `json:"awsAccountId"`
	ConfigurationItemStatus       string                     `json:"configurationItemStatus"`
	ResourceType                  string                     `json:"resourceType"`
	ResourceID                    string                     `json:"resourceId"`
	ResourceName                  string                     `json:"resourceName"`
	ARN                           string                     `json:"ARN"`
	AwsRegion                     string                     `json:"awsRegion"`
	AvailabilityZone              string                     `json:"availabilityZone"`
	ConfigurationStateMd5Hash     string                     `json:"configurationStateMd5Hash"`
	ResourceCreationTime          time.Time                  `json:"resourceCreationTime"`
	ConfigurationItemDeliveryTime interface{}                `json:"configurationItemDeliveryTime"`
	RecordingFrequency            interface{}                `json:"recordingFrequency"`
}

type Configuration struct {
	Name         string    `json:"name"`
	Owner        Owner     `json:"owner"`
	CreationDate time.Time `json:"creationDate"`
	Region       string    `json:"region"`
}

type Owner struct {
	DisplayName interface{} `json:"displayName"`
	ID          string      `json:"id"`
}

type SupplementaryConfiguration struct {
	AccessControlList                 string                            `json:"AccessControlList"`
	BucketAccelerateConfiguration     BucketAccelerateConfiguration     `json:"BucketAccelerateConfiguration"`
	BucketLoggingConfiguration        BucketLoggingConfiguration        `json:"BucketLoggingConfiguration"`
	BucketNotificationConfiguration   BucketNotificationConfiguration   `json:"BucketNotificationConfiguration"`
	BucketPolicy                      BucketPolicy                      `json:"BucketPolicy"`
	BucketTaggingConfiguration        BucketTaggingConfiguration        `json:"BucketTaggingConfiguration"`
	BucketVersioningConfiguration     BucketVersioningConfiguration     `json:"BucketVersioningConfiguration"`
	IsRequesterPaysEnabled            bool                              `json:"IsRequesterPaysEnabled"`
	PublicAccessBlockConfiguration    PublicAccessBlockConfiguration    `json:"PublicAccessBlockConfiguration"`
	ServerSideEncryptionConfiguration ServerSideEncryptionConfiguration `json:"ServerSideEncryptionConfiguration"`
}

type BucketAccelerateConfiguration struct {
	Status             interface{} `json:"status"`
	IsRequesterCharged bool        `json:"isRequesterCharged"`
}

type BucketLoggingConfiguration struct {
	DestinationBucketName interface{} `json:"destinationBucketName"`
	LogFilePrefix         interface{} `json:"logFilePrefix"`
	TargetObjectKeyFormat interface{} `json:"targetObjectKeyFormat"`
}

type BucketNotificationConfiguration struct {
	Configurations           Configurations `json:"configurations"`
	EventBridgeConfiguration interface{}    `json:"eventBridgeConfiguration"`
}

type Configurations struct {
}

type BucketPolicy struct {
	PolicyText interface{} `json:"policyText"`
}

type BucketTaggingConfiguration struct {
	TagSets []TagSet `json:"tagSets"`
}

type TagSet struct {
	Tags Tags `json:"tags"`
}

type Tags struct {
	Sampe string `json:"sampe"`
}

type BucketVersioningConfiguration struct {
	Status             string      `json:"status"`
	IsMfaDeleteEnabled interface{} `json:"isMfaDeleteEnabled"`
}

type PublicAccessBlockConfiguration struct {
	BlockPublicAcls       bool `json:"blockPublicAcls"`
	IgnorePublicAcls      bool `json:"ignorePublicAcls"`
	BlockPublicPolicy     bool `json:"blockPublicPolicy"`
	RestrictPublicBuckets bool `json:"restrictPublicBuckets"`
}

type ServerSideEncryptionConfiguration struct {
	Rules []Rule `json:"rules"`
}

type Rule struct {
	ApplyServerSideEncryptionByDefault ApplyServerSideEncryptionByDefault `json:"applyServerSideEncryptionByDefault"`
	BucketKeyEnabled                   bool                               `json:"bucketKeyEnabled"`
}

type ApplyServerSideEncryptionByDefault struct {
	SseAlgorithm   string      `json:"sseAlgorithm"`
	KmsMasterKeyID interface{} `json:"kmsMasterKeyID"`
}
