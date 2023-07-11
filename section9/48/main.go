package main

// スライス
// copy

import "fmt"

func main() {
		var sl []int = []int{100, 200}
		var sl2 []int = sl // slの要素をsl2にコピーする

		// 参照型の場合は参照渡しになるので、sl2の要素を変更するとslの要素も変更される
		// 参照型：スライス、マップ、チャネル
		sl2[0] = 1000 // sl2のインデックス番号0の要素に1000を代入。slの要素も変更される(参照型特有の性質)
		fmt.Println(sl) // [1000 200]

		// cf. 基本型の場合
		var i int = 100
		var i2 int = i // iの値をi2にコピーする

		i2 = 200 // i2に200を代入してもiの値は変わらない
		fmt.Println(i) // 100
		fmt.Println(i2) // 200

		// copy関数を使ってスライスをコピーする
		// copy(コピー先のスライス, コピー元のスライス)
		var sl3 []int = []int{1, 2, 3, 4, 5}
		var sl4 []int = make([]int, 5) // 要素数5のint型のスライスを作成
		fmt.Println(sl4)
		n := copy(sl4, sl3) // sl3の要素をsl4にコピーする。先頭からコピーする。nはコピーに成功した数

		fmt.Println(sl3) // [1 2 3 4 5]
		fmt.Println(sl4) // [1 2 3 4 5]
		fmt.Println(n) // 5

		// copy関数を使ってコピーしたので参照渡しは起きない
		sl4[0] = 1000 // sl4のインデックス番号0の要素に1000を代入
		fmt.Println(sl3) // [1 2 3 4 5]
		fmt.Println(sl4) // [1000 2 3 4 5]
}
