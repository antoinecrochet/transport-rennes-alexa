# Transport Rennes Alexa Skill

This project is an Alexa skill built with Go that provides information about transport services in Rennes. It uses AWS Lambda for serverless deployment and DynamoDB for data storage.

## Prerequisites

- Go 1.23 or later
- AWS CLI configured with appropriate permissions
- An Alexa Developer account
- AWS Lambda and DynamoDB setup

## Project Structure

- `cmd/main.go`: Entry point for the Lambda function.
- `internal/alexa-skill`: Contains the Alexa skill handlers and locale strings.
- `internal/dynamoDB`: Contains the DynamoDB client initialization and operations.

## Build

Run the following command to build the project:
```sh
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap cmd/main.go
```

## Deployment

Documentation about deployment on AWS is available here: https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html