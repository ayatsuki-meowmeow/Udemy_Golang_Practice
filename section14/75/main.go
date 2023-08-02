package main

// テスト
// テスト用のパッケージは標準パッケージにある
// テストをしたいファイルと同じ階層に、_test.goというファイルを作成する
// 今回だと、main.goと同じ階層に、main_test.goを作成する

import (
		"fmt"

		"golang_udemy/section14/75/alib"
)

func IsOne(i int) bool {
		if i == 1 {
				return true
		} else {
				return false
		}
}

func main() {
		fmt.Println(IsOne(1))
		fmt.Println(IsOne(2))

		// alibパッケージのAverage()を使う
		// Average()の引数に渡すスライスを作成する
		s := []int{1, 2, 3, 4, 5}
		// Average()の戻り値を表示する
		fmt.Println(alib.Average(s))
}
