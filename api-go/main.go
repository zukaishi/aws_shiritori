package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	list := "ミュウ,ウパー,パチリス"

	return events.APIGatewayProxyResponse{
		Body:       string(list),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type Response struct {
	List string `json:"list"`
}
