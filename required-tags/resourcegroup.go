package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/resourcegroupstaggingapi"
)

type ResourceGroupClient interface {
	TagResources(input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error)
}

type ResourceGroupRepository interface {
	TagResources(input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error)
}

type ResourceGroupRepositoryImpl struct {
	client ResourceGroupClient
}

func NewResourceGroupRepository(sess *session.Session) *ResourceGroupRepositoryImpl {
	client := resourcegroupstaggingapi.New(sess)
	return &ResourceGroupRepositoryImpl{client: client}
}

func (s *ResourceGroupRepositoryImpl) TagResources(input *resourcegroupstaggingapi.TagResourcesInput) (*resourcegroupstaggingapi.TagResourcesOutput, error) {
	return s.client.TagResources(input)
}
