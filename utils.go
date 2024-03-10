package utils

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

const (
	queueURL = "http://localhost:4566/000000000000/local-test-queue"
)

func GetSQSClient() *sqs.Client {

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "http://localhost:4566",
				SigningRegion: "us-east-1",
			}, nil
		})),
	)
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		return nil
	}

	return sqs.NewFromConfig(cfg)
}
