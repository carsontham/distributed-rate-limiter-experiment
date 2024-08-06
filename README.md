# Experimenting with rate limiting using Redis

To experiment with various open-source rate limiting algorithms, [LocalStack](https://github.com/localstack/localstack) SQS service will be used to mimic a local Amazon Simple Queue Service (SQS).

# Understanding rate limiting on distributed systems

When working with distributed systems, it is common to deploy multiple pods of the same service for high availability and performance. This experiment aims to demonstrate how various rate limiting algorithm works.

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
- Apart from code changes to introduce rate limiting, load balancers can be considered for this problem. However, this may be leaning towards infra-related changes and may be cumbersome should the rate-limit be changed.
