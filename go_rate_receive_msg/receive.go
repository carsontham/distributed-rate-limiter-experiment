package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
)

const (
	queueURL = "http://localhost:4566/000000000000/local-test-queue"
)

func main() {
	// Create a new SQS client
	sqsClient := GetSQSClient()

	// Create a new Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// allow user to choose whether to reset the db global counter
	var flushOption int
	fmt.Print("Flush DB? (-1 to flush): ")
	fmt.Scanln(&flushOption)

	// flush db if input == -1

	if flushOption == -1 {
		fmt.Println("Flushing DB...")
		_, err := rdb.FlushDB(context.Background()).Result()
		if err != nil {
			fmt.Println("Error clearing database:", err)
			return
		}

		err = rdb.Set(context.Background(), "global_counter", 0, 0).Err()
		if err != nil {
			fmt.Println("Error setting global counter:", err)
			return
		}
	}

	// Create a new rate limiter

	limiter := rate.NewLimiter(rate.Limit(1), 20)

	// Start multiple goroutines to receive messages
	// var wg sync.WaitGroup
	// for i := 0; i < 3; i++ {
	// 	wg.Add(1)
	// 	go func() {
	// 		receiveMessages(sqsClient, queueURL, limiter, rdb)
	// 		wg.Done()
	// 	}()
	// }

	// wg.Wait()

	receiveMessages(sqsClient, queueURL, limiter, rdb)
}
func receiveMessages(client *sqs.Client, queueURL string, limiter *rate.Limiter, rdb *redis.Client) {
	ctx := context.Background()

	for {
		//time.Sleep(1 * time.Second)

		// Get the global counter value from Redis
		counter, err := rdb.Get(ctx, "global_counter").Int()
		if err != nil {
			fmt.Println("Error getting global counter:", err)
			return
		}

		fmt.Println("Global counter:", counter)
		// If the counter has hit 500, stop receiving messages
		if counter >= 1000 {
			fmt.Println("Global counter has hit 1000. Stopping.")
			return
		}

		// Check if we're allowed to receive messages
		fmt.Println("Tokens left: ", limiter.Tokens())
		if limiter.Allow() {
			// fmt.Println("Tokens left: ", limiter.Tokens())
			fmt.Println("Rate limit not exceeded. messages as follows.")
			// Receive up to 10 messages from the queue
			result, err := client.ReceiveMessage(context.TODO(), &sqs.ReceiveMessageInput{
				QueueUrl:            &queueURL,
				MaxNumberOfMessages: 10,
			})
			if err != nil {
				fmt.Println("Error receiving messages:", err)
				return
			}

			// If no messages were returned, the queue is empty
			if len(result.Messages) == 0 {
				fmt.Println("Queue is now empty.")
			}

			// Process and delete received messages
			for _, message := range result.Messages {
				// Increment the global counter in Redis
				err := rdb.Incr(ctx, "global_counter").Err()
				if err != nil {
					fmt.Println("Error incrementing global counter:", err)
					return
				}

				fmt.Println("Body:", *message.Body)

				_, err = client.DeleteMessage(context.TODO(), &sqs.DeleteMessageInput{
					QueueUrl:      &queueURL,
					ReceiptHandle: message.ReceiptHandle,
				})
				if err != nil {
					fmt.Println("Error deleting message:", err)
					return
				}
			}

			fmt.Println("---------------------------------------------------")
		} else {
			fmt.Println("Tokens left: ", limiter.Tokens())
			fmt.Println("Rate limit exceeded. Waiting before receiving more messages.")
			time.Sleep(time.Second)
			continue
		}
	}
}

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
