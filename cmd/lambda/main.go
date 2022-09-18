package main

import (
	"context"
	"gostorefront/internal/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	response, _ := app.HandleRequest(app.Request{
		Body:   request.Body,
		Path:   request.RequestContext.HTTP.Path,
		Method: request.RequestContext.HTTP.Method,
	})
	headers := map[string]string{"Content-Type": "text/html; charset=utf-8"}
	return events.LambdaFunctionURLResponse{Body: response.Body, StatusCode: response.StatusCode, Headers: headers}, nil
}

func main() {
	lambda.Start(handler)
}
