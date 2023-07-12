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
		m4 := make(map[int]string) // キーが文字列、値が整数のマップ
		fmt.Println(m4) // map[]
		fmt.Println(len(m4)) // 0

		// 要素の追加
		// 変数名[キー] = 値
		m4[1] = "JAPAN"
		m4[2] = "USA"
		fmt.Println(m4) // map[1:JAPAN 2:USA]

		// 要素の取得
		// 変数名[キー]
		fmt.Println(m["A"]) // 100
		fmt.Println(m4[1]) // JAPAN

		// 登録されていないキーを指定した場合はゼロ値(=初期値)が返る
		// nilのようなものは返ってこない
		fmt.Println(m["C"]) // 0
		fmt.Println(m4[3]) // ""

		// 登録されていないキーを指定した場合にエラーが返るようにハンドリングする
		// 取り出しに成功したら、okにtrueが入る
		s, ok := m4[1]
		fmt.Println(s, ok) // JAPAN true
		// 取り出しに失敗したら、okにfalseが入る
		s2, ok := m4[3]
		if !ok {
				fmt.Println("error") // error
		}
		fmt.Println(s2, ok) //  false
		// エラーハンドリングだけをするときは、_を使って変数を捨てることができる
		_, ok2 := m4[3]
		if !ok2 {
				fmt.Println("error") // error
		}

		m4[2] = "CHINA"
		fmt.Println(m4) // map[1:JAPAN 2:CHINA]

		m4[3] = "US"
		fmt.Println(m4) // map[1:JAPAN 2:CHINA 3:US]

		// 要素の削除
		// delete(変数名, キー)
		delete(m4, 3)
		fmt.Println(m4) // map[1:JAPAN 2:CHINA]

		// 要素数の取得
		// len(変数名)
		fmt.Println(len(m4)) // 2
}
