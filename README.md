Steps to Set Up LocalStack: https://docs.localstack.cloud/user-guide/aws/sqs/ 

Pull LocalStack Docker Image:

Open a terminal and run the following command to pull the LocalStack Docker image:
docker pull localstack/localstack

Run LocalStack Container:
Once the image is downloaded, run the LocalStack container using the following command:

docker run -p 4566:4566 -p 8080:8080 -e SERVICES=sqs localstack/localstack
This command maps the ports 4566 (AWS services) and 8080 (LocalStack web UI) from the container to your local machine. The SERVICES=sqs environment variable specifies that only the SQS service should be started.

Wait for LocalStack to Start:
LocalStack may take a moment to start. Once it's ready, you should see log messages indicating that services are available.
Access LocalStack Web UI:

Open your web browser and navigate to http://localhost:8080 to access the LocalStack web UI. This interface provides insights into the services and their status.
Configure AWS CLI:

Configure your AWS CLI to interact with LocalStack. Run the following command and enter placeholder values for access key, secret key, region, and output format:


aws configure --profile localstack
AWS Access Key ID: Enter any value (e.g., test).
AWS Secret Access Key: Enter any value (e.g., test).
Default region name: Enter us-east-1.
Default output format: Enter json.



Test SQS in LocalStack:
Now that LocalStack is set up, let's test SQS:

Create an SQS Queue:

Run the following command to create an SQS queue using the AWS CLI and the localstack profile:

bash
Copy code
aws sqs create-queue --queue-name local-test-queue --profile localstack
Note the QueueUrl in the response; you'll use it in subsequent commands.

Send a test message to the created SQS queue:
http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/local-test-queue
aws sqs send-message --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/local-test-queue --message-body "Hello, LocalStack!" --profile localstack
Replace <Your-Queue-URL> with the actual QueueUrl from the previous step.

Receive messages from the SQS queue:

aws sqs receive-message --queue-url http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/local-test-queue --max-number-of-messages 1 --profile localstack
Again, replace <Your-Queue-URL> with the actual QueueUrl.


In the LocalStack response you provided, it looks like the CreateQueue operation was successful, and the status code 200 indicates a successful request. However, it seems that the QueueUrl wasn't directly printed in the log. You can obtain the QueueUrl separately by using the aws sqs get-queue-url command. Here's how you can do it:

aws sqs get-queue-url --queue-name local-test-queue --profile localstack
This command retrieves the URL of the SQS queue named local-test-queue. The response will include the QueueUrl that you can use for sending and receiving messages.

aws sqs send-message --queue-url <Your-Queue-URL> --message-body "Hello, LocalStack!" --profile localstack
Replace <Your-Queue-URL> with the actual QueueUrl obtained from the get-queue-url command.

aws sqs get-queue-url --queue-name local-test-queue --profile localstack
This command should return a JSON response containing the QueueUrl. Ensure that the QueueUrl is correctly printed.

