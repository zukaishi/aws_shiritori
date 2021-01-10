package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, params interface{}) (interface{}, error) {
	// TODO your processing
	return params, nil
}

func main() {
	lambda.Start(HandleRequest)
}
