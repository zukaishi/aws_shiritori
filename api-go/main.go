package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name1 := request.QueryStringParameters["name1"]
	name2 := request.QueryStringParameters["name2"]
	old := 25

	person := PersonResponse{
		Name1: name1,
		Name2: name2,
		Old:   old,
	}
	jsonBytes, _ := json.Marshal(person)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

type PersonResponse struct {
	Name1 string `json:"name1"`
	Name2 string `json:"name2"`
	Old   int    `json:"old"`
}
