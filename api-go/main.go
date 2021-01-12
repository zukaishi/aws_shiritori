package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	list := "ミュウ,ウパー,パチリス"

	result := Response{
		List: list,
	}
	jsonBytes, _ := json.Marshal(result)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type Response struct {
	List string `json:"list"`
}
