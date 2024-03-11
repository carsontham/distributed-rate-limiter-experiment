# Distributed Rate Limiting with Redis
This experiment aims to demonstrate how to achieve distributed rate limiting using Redis :)

To mimic a real-life problem that can occur in organizations, this project uses [LocalStack](https://github.com/localstack/localstack) SQS service which mimics a local Amazon Simple Queue Service (SQS).

# Simulated Real Life Problem
For example, you have an application subscribing to a Amazon SQS which continuously polls for messages. For each message polled, the application wants to perform certain actions, such as sending an API request to an external organization. However, if there are large amount of messages in the queue, it would mean large amount of requests being sent. However, external organizations does not want to be flooded with too many incoming API request. Thus, rate limiting is introduced. 

But if the application runs on multiple pods/servers, simple rate-limiting may not be sufficient since each pod/server would only know information about itself and not information of other pods. This is where Distributed Rate Limiting can be achieved using Redis.

# Technologies used to experiment with Distributed Rate Limiting 
- Using LocalStack to mimic AWS services https://github.com/localstack/localstack 
- Redis client
- Using different open-source rate-limiting tools:
    - rate https://pkg.go.dev/golang.org/x/time/rate (Token Bucket Algorithm)
    - redis-rate https://github.com/go-redis/redis_rate (Leaky Bucket Algorithm)


# Prerequisite
- Set up localstack and initialize the SQS service using the following command:
    
    `docker run -p 4566:4566 -p 8080:8080 -e SERVICES=sqs localstack/localstack`

- Verify that the LocalStack SQS is running:

    `localstack status services`

- Install Redis and initialize Redis server
    
    `redis-server`

- To populate the queue, send messages:
    `go run send.go`

- Receiving messages from distributed rate limiter using go-rate package using Redis:
  
    `go run receive.go`

- Receiving messages from distributed rate limiter using redis-rate package using Redis:
  
    `go run receive.go`

# Learning Outcome
- A Distributed Rate Limiter can be achieved by pairing Redis together with any of the existing rate-limiter packages
- To prevent potential race conditions, locks can be introduced to ensure that the rate limiters running on different pods are always accessing accurate global counter values
     
