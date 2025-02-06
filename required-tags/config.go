package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"
)

type ConfigClient interface {
	PutEvaluationsWithContext(ctx aws.Context, input *configservice.PutEvaluationsInput, opts ...request.Option) (*configservice.PutEvaluationsOutput, error)
}

type ConfigRepository interface {
	PutEvaluations(ctx context.Context, input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error)
}

type ConfigRepositoryImpl struct {
	client ConfigClient
}

// NewConfigRepository creates a new instance of ConfigRepositoryImpl
func NewConfigRepository(sess *session.Session) *ConfigRepositoryImpl {
	client := configservice.New(sess)
	return &ConfigRepositoryImpl{client: client}
}

// PutEvaluations sends evaluation results to AWS Config
func (s *ConfigRepositoryImpl) PutEvaluations(ctx context.Context, input *configservice.PutEvaluationsInput) (*configservice.PutEvaluationsOutput, error) {
	return s.client.PutEvaluationsWithContext(ctx, input)
}
