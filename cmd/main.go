package main

import (
	"context"
	"fmt"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	u, err := url.Parse(request.Headers["referer"])
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Path)

	return events.LambdaFunctionURLResponse{Body: request.Body, StatusCode: 200}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
