package main

import "fmt"

// 文字列型
// string

func main() {
		var s string = "Hello Golang"
		fmt.Println(s)
		fmt.Printf("%T\n", s) // string

		// var s2 string = 'Hello'
		// シングルクォーテーションで囲うとエラーになる

		var si string = "300"
		fmt.Println(si)
		fmt.Printf("%T\n", si)

		// 複数行に渡る表示にはバッククォーテーションを使う
		fmt.Println(`test
		test
				test`)

		// ダブルクォーテーションを文字列として扱うにはバックスラッシュを入れる
		fmt.Println("aaa\"aaa")

		// バッククォーテーションの中でダブルクォーテーションを使っても、文字列として扱える
		var s2 string = `"aaaa"`
		fmt.Println(s2)

		// 文字列はバイト配列の集まり
		// バイトについては後の講義でやる
		fmt.Println(s[0]) // 72
		fmt.Print(string(s[0])) // H
}
