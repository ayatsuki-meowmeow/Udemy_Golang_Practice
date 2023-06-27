package main

// if
// 条件分岐
// エラーハンドリング

import (
	"fmt"
	"strconv"
)

func main() {
		var s string = "A"

		i, err := strconv.Atoi(s)
		// エラーが発生した場合、errにエラー内容が入る
		if err != nil {
				fmt.Println(err)
		}
		fmt.Printf("i = %T\n", i)
}
