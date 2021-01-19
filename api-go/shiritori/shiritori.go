package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	name1 := "ゼニガメ"
	name2 := "スバメ"

	pokemonList := map[int]string{
		7:   "ゼニガメ",
		132: "メタモン",
		276: "スバメ",
		376: "メタグロス",
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
	fmt.Println(result)
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
