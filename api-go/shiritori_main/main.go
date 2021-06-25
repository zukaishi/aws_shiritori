package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name1 := request.QueryStringParameters["name1"]
	name2 := request.QueryStringParameters["name2"]
	list := shiritori(name1, name2)
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

func shiritori(name1 string, name2 string) string {
	url := "http://aws-shiritori.tk/data/pokemon_list.csv"
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
		if name2 == record[1] || getLastString(record[1]) != "ン" {
			var i int
			i, _ = strconv.Atoi(record[0])
			pokemonList[i] = record[1]
		}
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
		result += fmt.Sprintf("%d:%s", no, pokemonList[no])
		word = pokemonList[no]
		delete(pokemonList, no)

		lastString := getLastString(word)
		list := containsList(pokemonList, lastString)
		if len(list) == 0 || endNo == no {
			break
		}
		result += ","
		if lastString == getRuneAt(name2, 0) {
			result += fmt.Sprintf("%d:%s", endNo, name2)
			break
		}

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

func getLastString(word string) string {
	// 対象文字の最後の文字を取得utf-8のため、/3している
	lastString := getRuneAt(word, len(word)/3-1)

	// 最後の文字が伸ばし棒の場合
	if lastString == "ー" {
		lastString = getRuneAt(word, len(word)/3-2)
	}

	// 特殊文字、捨て仮名を扱いやすい形へ変換する
	r := strings.NewReplacer("♂", "ス", "♀", "ス", "ァ", "ア", "ィ", "イ", "ゥ", "ウ", "ェ", "エ", "ォ", "オ", "ュ", "ユ", "ャ", "ヤ", "ョ", "ヨ")
	resStr := r.Replace(lastString)
	return resStr
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
