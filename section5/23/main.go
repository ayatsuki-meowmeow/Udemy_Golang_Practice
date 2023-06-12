package main

import (
		"fmt"
)

// 定数
// 後から値を上書きすることができない
// 後から値を変えることがないものはセキュリティ上定数にする方がいい
// 定数はグローバルに書くことが多いので、関数外に書いて他のパッケージからも呼び出すことができるようにする

// constを使う
// 定数名の頭文字を大文字にすることでグローバルになる

const Pi = 3.14

// 複数の定数の定義
const (
		URL = "https://xxx.co.jp"
		SiteName = "test"
)

// 値の省略
// B,CにはAの値が、E,FにはDの値が入る
const (
		A = 1
		B
		C
		D = "A"
		E
		F
)

// 定数には原則最大値がない
// var Big int = 9223372036854775807 + 1 // int型の最大値+1 エラーになる
const Big = 9223372036854775807 + 1 // エラーにならない

// iotaは連続する整数を生成する
const (
		c0 = iota
		c1
		c2
)

func main() {
		fmt.Println(Pi) // 3.14

		// Pi = 3
		// fmt.Println(Pi)
		// ↑定数は上書きできないからエラーになる

		fmt.Println(URL, SiteName)

		fmt.Println(A, B, C, D, E, F) // 1 1 1 A A A

		fmt.Println(Big - 1)

		fmt.Println(c0, c1, c2) // 0 1 2
}
