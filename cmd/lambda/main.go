package main

import (
	"context"
	"lambdacommerce/internal/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	response, _ := app.HandleRequest(app.Request{
		Body:   request.Body,
		Path:   request.RequestContext.HTTP.Path,
		Method: request.RequestContext.HTTP.Method,
	})
	return events.LambdaFunctionURLResponse{Body: response.Body, StatusCode: response.StatusCode}, nil
}

func main() {
	lambda.Start(handler)
}
