package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name1 := request.QueryStringParameters["name1"]
	list := comprised(name1)
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

func comprised(word string) string {
	url := "https://s3-ap-northeast-1.amazonaws.com/website.shiritori.com/data/pokemon_list.csv"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	pokemonList := map[int]string{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var i int
		i, _ = strconv.Atoi(record[0])
		pokemonList[i] = record[1]
	}

	result := ""
	for key, value := range pokemonList {
		if strings.Contains(value, word) {
			result += fmt.Sprintf("%d:%s", key, value)
			result += ","
		}
	}
	return result
}
