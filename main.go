package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/herbal-goodness/inventoryflo-api/pkg/util/db"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/herbal-goodness/inventoryflo-api/pkg/router"
)

func main() {
	lambda.Start(handler)
	defer db.CloseDb()
}

func handler(_ context.Context, evt events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	method := evt.HTTPMethod
	if method == "OPTIONS" {
		allowed := "OPTIONS, GET, POST, PUT, DELETE"
		optionsHeaders := map[string]string{
			"Allow":                        allowed,
			"Access-Control-Allow-Methods": allowed,
			"Content-Length":               "0",
		}
		return respond(200, "", optionsHeaders)
	}

	path := strings.Split(evt.Path, "/")[1:]

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(evt.Body), &body); err != nil {
		errMsg := fmt.Sprintf("Unable to unmarshal request body to json: %v", err)
		return respond(500, errMsg, nil)
	}

	resp, err := router.Route(method, path, body)
	if err != nil {
		return respond(500, err.Error(), nil)
	}

	response, err := json.Marshal(resp)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to marshal response body to string: %v", err)
		return respond(500, errMsg, nil)
	}

	return respond(200, string(response), nil)
}

func respond(status int, body string, headers map[string]string) (events.APIGatewayProxyResponse, error) {
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
	}, nil
}
