package main

import (
	"context"
	"gostorefront/pkg/request"
	"gostorefront/routes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, functionURLRequest events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	response := routes.Handle(
		request.Request{
			Body:   functionURLRequest.Body,
			Path:   functionURLRequest.RequestContext.HTTP.Path,
			Method: functionURLRequest.RequestContext.HTTP.Method,
		},
	)
	return events.LambdaFunctionURLResponse{Body: response.Body, StatusCode: response.StatusCode, Headers: response.Headers}, nil
}

func main() {
	lambda.Start(handler)
}
