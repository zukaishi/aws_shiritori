package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("game start \n")

	// 2. pokemonList[1] -> "フシギダネ"　配列で「全体ポケモンリスト」へ格納する。
	pokemonList := map[int]string{
		7:   "ゼニガメ",
		132: "メタモン",
		276: "スバメ",
		376: "メタグロス",
	}

	// 3. 入力された二匹のポケモンの名前が存在しているかチェック
	name1 := "ゼニガメ"
	name2 := "スバメ"

	no1 := contains(pokemonList, name1)
	if no1 == 0 {
		fmt.Printf("end1 ¥n")
	}
	no2 := contains(pokemonList, name2)
	if no2 == 0 {
		fmt.Printf("end2 ¥n")
	}
	fmt.Printf("check complete \n")

	// 4. ループ開始
	word := ""
	no := no1
	resultList := map[int]string{
		no1: name1,
	}

	for i, v := range pokemonList {
		fmt.Println(i, v)

		word = pokemonList[no]
		delete(pokemonList, no)

		// 対象文字の最後の文字を取得utf-8のため、/3している
		lastString := getRuneAt(word, len(word)/3-1)
		list := containsList(pokemonList, lastString)

		if len(list) == 0 {
			break
		}

		randMap := []int{}
		for key := range list {
			randMap = append(randMap, key)
		}
		rand.Seed(time.Now().UnixNano())
		randNo := rand.Intn(len(randMap))
		no = randMap[randNo]
		resultList[no] = pokemonList[no]
		fmt.Println("no=", no)
	}

	fmt.Println("result")
	result := ""
	for i, v := range resultList {
		result += v + ","
		fmt.Println(i, v)
	}
	fmt.Println(result)
	fmt.Println("end")
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
