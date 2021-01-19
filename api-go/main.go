package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	list := shiritori()
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

func shiritori() string {
	name1 := "ゼニガメ"
	name2 := "スバメ"

	pokemonList := map[int]string{}
	url := "https://s3-ap-northeast-1.amazonaws.com/website.shiritori.com/data/pokemon_list.csv"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	for {
		// 1行づつ読み込む
		record, err := reader.Read()
		// ファイルの末尾で処理終了
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

	startNo := contains(pokemonList, name1)
	endNo := contains(pokemonList, name2)
	if startNo == 0 || endNo == 0 {
		fmt.Printf("input name failed ¥n")
	}

	word := ""
	no := startNo
	result := ""
	for {
		result += pokemonList[no]
		word = pokemonList[no]
		delete(pokemonList, no)

		// 対象文字の最後の文字を取得utf-8のため、/3している
		lastString := getRuneAt(word, len(word)/3-1)
		list := containsList(pokemonList, lastString)
		if len(list) == 0 || endNo == no {
			break
		}
		result += ","

		randMap := []int{}
		for key := range list {
			randMap = append(randMap, key)
		}
		rand.Seed(time.Now().UnixNano())
		randNo := rand.Intn(len(randMap))
		no = randMap[randNo]
	}
	return result
}
func contains(pokemonList map[int]string, name string) int {
	for i := range pokemonList {
		if pokemonList[i] == name {
			return i
		}
	}
	return 0
}

func containsList(pokemonList map[int]string, str string) map[int]string {
	list := map[int]string{}
	for i := range pokemonList {
		if getRuneAt(pokemonList[i], 0) == str {
			list[i] = pokemonList[i]
		}
	}
	return list
}

func getRuneAt(s string, i int) string {
	rs := []rune(s)
	return string(rs[i])
}
