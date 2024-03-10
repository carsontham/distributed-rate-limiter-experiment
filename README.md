# Distributed Rate Limiting with Redis
This experiment aims learn how to achieve distributed rate limiting using Redis :)

To mimic a real-life problem that can occur in organizations, this project uses LocalStack SQS service which mimics a local Amazon Simple Queue Service (SQS).

# Stimulated Problem
For example, you have an application subscribing to a AWS Queue which continuously polls for messages. For each message polled, the application wants to send a request to an external organization. However, external organization does not want to be floaded with too many incoming API request. Thus, rate limiting is introduced. 

But if the application runs on multiple pods/servers, simple rate-limiting may not be sufficient since each pod only know information about itself and not information of other pods. This is where Distributed Rate Limiting can be achieved using Redis :D

# Technologies used to experiment with Distributed Rate Limiting 
- Using LocalStack to mimic AWS services https://github.com/localstack/localstack 
- Redis client
- Using different open-source rate-limiting tools:
    - rate https://pkg.go.dev/golang.org/x/time/rate
    - redis-rate https://github.com/go-redis/redis_rate
    - uber-go https://github.com/uber-go/ratelimit


# Pre-requisite
- Set up localstack and initialize it using following commands:
    
    `docker run -p 4566:4566 -p 8080:8080 -e SERVICES=sqs localstack/localstack`

- Verify it is the LocalStack SQS is running via 

    `localstack status services`

- Install Redis and initialize redis server
    
    `redis-server`

- To populate the queue, send messages via:
    `go run send.go`

- Receiving messages from distributed rate limiter using go-rate package using Redis:
  
    `go run receive.go`

- Receiving messages from distributed rate limiter using redis-rate package using Redis:
  
    `go run receive.go`
