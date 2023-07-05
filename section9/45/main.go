package main

import "fmt"

// スライス
// 宣言、操作
// 配列によく似たデータ型

func main() {
		var sl []int // int型のスライスの宣言 []に要素数を指定しない
		fmt.Println(sl) // []

		var sl2 []int = []int{100, 200} // 明示的な宣言で値を格納
		fmt.Println(sl2) // [100 200]

		sl3 := []string{"A", "B"} // 暗黙的な宣言で値を格納
		fmt.Println(sl3) // [A B]

		// make関数を使ってスライスを作成する
		// make(スライスの型, スライスの長さ(要素数), スライスの容量)
		sl4 := make([]int, 5) // 要素数が5のint型のスライスを作成 int型の初期値は0なので、[0 0 0 0 0]となる
		fmt.Println(sl4) // [0 0 0 0 0]

		// 値の代入
		sl2[0] = 1000 // sl2のインデックス番号0の要素に1000を代入
		fmt.Println(sl2) // [1000 200]

		// 値の取り出し
		sl5 := []int{1, 2, 3, 4, 5}
		fmt.Println(sl5) // [1 2 3 4 5]
		fmt.Println(sl5[0]) // 1
		fmt.Println(sl5[2:4]) // [3 4] インデックス番号2から4の手前を取得
		fmt.Println(sl5[:2]) // [1 2] インデックス番号0から2の手前を取得
		fmt.Println(sl5[2:]) // [3 4 5] インデックス番号2から最後までを取得
		fmt.Println(sl5[:]) // [1 2 3 4 5] インデックス番号0から最後までを取得
		fmt.Println(sl5[len(sl5) - 1]) // 5 len関数でスライスの要素数を取得し、-1することで最後の要素を取得できる
		fmt.Println(sl5[1:len(sl5) - 1]) // [2 3 4] スライスの最初の要素と最後の要素を除いた要素を取得する
}
