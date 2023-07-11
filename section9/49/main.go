package main

// スライス
// for

import "fmt"

func main() {
		var sl []string = []string{"A", "B", "C"}
		fmt.Println(sl) // [A B C]

		// for文でスライスの要素を取り出す
		// for 変数名1, 変数名2 := range スライス名 {}
		// 変数名1はインデックス番号、変数名2は要素の値
		// rangeでスライスの長さ分だけ繰り返す
		for i, v := range sl {
				fmt.Println(i, v) // 0 A, 1 B, 2 C
		}

		// インデックス番号は不要な場合は_で省略できる
		for _, v := range sl {
				fmt.Println(v) // A, B, C
		}

		// インデックス番号のみ必要な場合は変数名2を省略できる
		for i := range sl {
				fmt.Println(i) // 0, 1, 2
		}

		// 変数2を省略しつつ、要素の値も取り出す方法
		for i := range sl {
				fmt.Println(i, sl[i]) // 0 A, 1 B, 2 C
		}

		// 古典的forを使ってスライスの要素を取り出す
		// for 変数名 := 0; 変数名 < len(スライス名); 変数名++ {}
		for i := 0; i < len(sl); i++ {
				fmt.Println(sl[i]) // A, B, C
		}
}
