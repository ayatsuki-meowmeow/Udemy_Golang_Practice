package main

import "fmt"

// マップ
// キーと値の組み合わせでデータを管理する

func main() {
		// マップの明示的宣言
		// var 変数名 map[キーの型]値の型 = map[キーの型]値の型{キー: 値, キー: 値, ...}
		var m map[string]int = map[string]int{"A": 100, "B": 200} // キーが文字列、値が整数のマップ
		fmt.Println(m) // map[A:100 B:200]

		// mapの暗黙的宣言
		// 変数名 := map[キーの型]値の型{キー: 値, キー: 値, ...}
		m2 := map[string]int{"A": 100, "B": 200} // キーが文字列、値が整数のマップ
		fmt.Println(m2) // map[A:100 B:200]

		// 改行して行末にカンマをつけることでも宣言できる
		// 改行して行末にカンマがある場合は最後の要素の後ろにもカンマが必要
		m3 := map[int]string{
				1: "A",
				2: "B",
		}
		fmt.Println(m3) // map[1:A 2:B]

		// make関数を使ってmapを作成する
		// make(map[キーの型]値の型, 要素数)
		// 要素数は省略可能で、省略した場合は0が設定される
		m4 := make(map[int]string) // キーが文字列、値が整数のマップで要素数は10
		fmt.Println(m4) // map[]

}
