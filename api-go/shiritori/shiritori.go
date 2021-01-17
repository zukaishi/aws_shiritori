package main

import (
	"fmt"
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
	for i, v := range pokemonList {
		if word == "" {
			word = pokemonList[no]
		}
		delete(pokemonList, no)

		// 対象文字の最後の文字を取得utf-8のため、/3している
		lastString := getRuneAt(word, len(word)/3-1)

		fmt.Println(lastString)

		fmt.Println(i, v)
	}
}

func contains(pokemonList map[int]string, name string) int {
	for i := range pokemonList {
		if pokemonList[i] == name {
			return i
		}
	}
	return 0
}

// func containsList(pokemonList map[int]string, str string, mode int) [] {
// 	let list = []
// 	for i := range pokemonList {
// 		if getRuneAt(pokemonList[i], 0) == str {
// 			list := append(list, i)
// 		}
// 	}
// 	return list
// }

func getRuneAt(s string, i int) string {
	rs := []rune(s)
	return string(rs[i])
}
