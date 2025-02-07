# Transport Rennes Alexa Skill

This project is an Alexa skill built with Go that provides information about transport services in Rennes. It uses AWS Lambda for serverless deployment and DynamoDB for data storage.

## Prerequisites

- Go 1.23 or later
- AWS CLI configured with appropriate permissions
- An Alexa Developer account
- AWS Lambda and DynamoDB setup

## Project Structure

- `cmd/lambda/main.go`: Entry point for the Lambda function.
- `internal/alexa-skill`: Contains the Alexa skill handlers and locale strings.
- `internal/dynamoDB`: Contains the DynamoDB client initialization and operations.
