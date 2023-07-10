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
}
