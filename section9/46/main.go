package main

// スライス
// append make len cap

import "fmt"

func main() {
		sl := []int{100, 200}
		fmt.Println(sl) // [100 200]

		// append関数を使ってスライスに要素を追加する
		// 配列は要素数が決まっているが、スライスは要素数が決まっていないので、要素を追加することができる
		// append(スライス, 追加する要素)
		// スライスの最後に追加する
		sl = append(sl, 300)
		fmt.Println(sl) // [100 200 300]

		// スライスの最後に複数の要素を追加する
		// append(スライス, 追加する要素1, 追加する要素2, ...)
		sl = append(sl, 400, 500, 600)
		fmt.Println(sl) // [100 200 300 400 500 600]

		// make関数を使ってスライスを作成する
		// make(スライスの型, スライスの長さ(要素数), スライスの容量)
		sl2 := make([]int, 5)
		fmt.Println(sl2) // [0 0 0 0 0]

		// len関数でスライスの要素数を取得する
		fmt.Println(len(sl2)) // 5

		// cap関数でスライスの容量を取得する
		fmt.Println(cap(sl2)) // 5

		// 容量10のスライスを作成する
		// プログラムのメモリーを気にする開発を行う場合は、容量を指定してスライスを作成する
		sl3 := make([]int, 5, 10)
		fmt.Println(sl3) // [0 0 0 0 0]
		fmt.Println(len(sl3)) // 5
		fmt.Println(cap(sl3)) // 10

		// 容量を超えた要素を追加する
		// 容量を超えた要素を追加すると、容量が自動的に倍増される
		// 倍増するからパフォーマンスへの影響がある
		sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)
		fmt.Println(sl3) // [0 0 0 0 0 1 2 3 4 5 6 7]
		fmt.Println(len(sl3)) // 12
		fmt.Println(cap(sl3)) // 20
}
