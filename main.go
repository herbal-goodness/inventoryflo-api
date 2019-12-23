package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(_ context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return response(200, req.Body, nil), nil
}

func response(status int, body string, headers map[string]string) events.APIGatewayProxyResponse {
	defaultHeaders := map[string]string{
		"Content-Type":                "application/json",
		"Access-Control-Allow-Origin": "*",
	}
	for k, v := range headers {
		defaultHeaders[k] = v
	}
	return events.APIGatewayProxyResponse{
		Body:            body,
		StatusCode:      status,
		Headers:         defaultHeaders,
		IsBase64Encoded: false,
	}
}
