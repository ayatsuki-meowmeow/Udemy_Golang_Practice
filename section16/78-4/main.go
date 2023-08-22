package main

import (
		"fmt"
		"time"
)

// time
// スリープ

func main() {
		// 指定時間のスリープ
		// time.Sleep()で指定した時間だけ実行中のゴルーチンをスリープする
		// time.Sleep(時間)

		// time.Duration型を引数に使うことで、スリープする時間をわかりやすく表現できる
		// 5 * time.Second の型を確認する
		fmt.Printf("%T\n", 5 * time.Second) // time.Duration

		// 5秒スリープする
		time.Sleep(5 * time.Second)
		fmt.Println("5秒経過しました")
}
