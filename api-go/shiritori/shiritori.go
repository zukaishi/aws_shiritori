package main

import (
	"fmt"
)

func main() {
	fmt.Printf("game start \n")

	/*
		flow & role
			 1. データは 1,フギダネ¥n 2,フシギバナ　...とは言ってくる想定
			 2. pokemonList[1] -> "フシギダネ"　配列で「全体ポケモンリスト」へ格納する。
			 3. 入力された二匹のポケモンの名前が存在しているかチェック
			 4. ループ開始
			 5. 	対象ポケモンの最後の文字を取り出す
			 6. 	最後の文字が伸ばし棒なら一つ前の文字を、次の最初の文字とする
			 7. 	最初の文字が決まったら、その文字から始まるポケモンを、「全体ポケモンリスト」から取得する
			 8. 	二匹目のポケモンが存在するかチェックし、存在していたら「結果ポケモンリスト」に詰め込み。ループ終了
			 9. 	なければ「ん」で終わるものを除外し、ランダムでその中から一匹を選択、「結果ポケモンリスト」に詰め込む、「全体ポケモンリスト」から除外する
			10. 	「ん」で終わるものしか存在しなければ、「結果ポケモンリスト」に詰め込み。ループ終了。「結果ポケモンリスト」がからならそのままループ終了
			10. ループ終了
			11. 「結果ポケモンリスト」を返却する

		ai.
			1. 10回実施して二匹目のポケモンに到達しなければ終了
			2. 実行した結果を保存、同じパターンのものは再度保存はしない
			3. 最短の経路を保存し返却できるようにしておく

	*/

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
	if !contains(pokemonList, name1) {
		fmt.Printf("end1 ¥n")
	}
	if !contains(pokemonList, name2) {
		fmt.Printf("end2 ¥n")
	}
	fmt.Printf("check end! \n")
}

func contains(pokemonList map[int]string, name string) bool {
	for i := range pokemonList {
		fmt.Printf("%s == %s \n", pokemonList[i], name)
		if pokemonList[i] == name {
			fmt.Printf("check! \n")
			return true
		}
	}
	return false
}
