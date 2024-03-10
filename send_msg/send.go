package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func main() {
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
		return
	}

	client := sqs.NewFromConfig(cfg)

	// Specify the URL of your SQS queue
	queueURL := "http://localhost:4566/000000000000/local-test-queue"

	// Number of messages to send
	numMessages := 10000

	// Messages to send
	for i := 1; i < numMessages+1; i++ {
		message := fmt.Sprintf("Message %d", i)

		// Send message
		_, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
			QueueUrl:    &queueURL,
			MessageBody: &message,
		})

		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		fmt.Printf("Sent message %d\n", i)
	}

	fmt.Println("All messages sent.")
}
