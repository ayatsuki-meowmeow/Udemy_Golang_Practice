package main

// マップ
// for

import "fmt"

func main() {
		var m map[string]int = map[string]int{
				"Apple": 100,
				"Banana": 200,
		} // キーが文字列、値が整数のマップ

		// for文でマップの要素を一つずつ取り出す
		// for 変数名, 変数名 := range 変数名 {}
		for k, v := range m {
				fmt.Println(k, v, "円")
				// フォーマット指定子 %s を使用してキーの文字列を表示し、%d を使用して整数の値を表示
				fmt.Printf("%s %d円\n", k, v)
		}

		// keyを使わない場合は_で省略できる
		for _, v := range m {
				fmt.Println(v)
		}

		// valueを使わない場合は宣言する必要がない
		for k := range m {
				fmt.Println(k)
		}
}

// %s: 文字列を表示するためのフォーマット指定子です。引数として文字列を渡すと、その文字列が表示されます。
// %d: 10進数の整数を表示するためのフォーマット指定子です。引数として整数を渡すと、その整数が表示されます。
