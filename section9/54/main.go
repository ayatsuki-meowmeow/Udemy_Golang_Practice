package main

// channel
// 複数のgoroutine間でデータの受け渡しをするために設計されたデータ構造
// データの送受信を行うデータ構造
// 送信専用、受信専用のチャネルを作成することも可能
// チャネルはキューのような構造になっている
// キューは先入れ先出し(FIFO)のデータ構造
// 宣言、操作

import "fmt"

func main() {
		// チャネルの宣言
		// 双方向のチャネル
		// var 変数名 chan データ型
		var ch1 chan int

		// 受信専用のチャネル
		// var 変数名 <-chan データ型
		// var ch2 <-chan int

		// 送信専用のチャネル
		// var 変数名 chan<- データ型
		// var ch3 chan<- int

		// make関数でチャネルとしての機能を持たせる
		// チャネルの生成と初期化を行って、チャネルを使用できるようにする
		// make(chan データ型)
		ch1 = make(chan int)
		// 直接make関数でチャネルを宣言することも可能
		ch2 := make(chan int)

		// チャネルには容量(バッファーサイズ)がある
		fmt.Println(cap(ch1)) // 0
		fmt.Println(cap(ch2)) // 0

		// バッファーサイズを指定してチャネルを作成する
		// make(chan データ型, バッファーサイズ)
		ch3 := make(chan int, 5)
		fmt.Println(cap(ch3)) // 5

		// データの操作
		// チャネルからデータを送信したり、受信したりすることができる

		// データの送信
		// 変数名 <- データ
		ch3 <- 1 // ch3というチャネルに1を送信する
		// len関数でチャネルが保持しているデータの数を取得できる
		fmt.Println(len(ch3)) // 1

		ch3 <- 2
		ch3 <- 3
		fmt.Println("len", len(ch3)) // 3

		// データの受信
		// 変数名 := <-チャネル
		i := <-ch3 // ch3というチャネルからデータを受信して、iに代入する
		fmt.Println(i) // 1 最初に入れたデータが取り出される
		fmt.Println("len", len(ch3)) // 2 次に入れたデータが取り出される

		var i2 int = <-ch3
		fmt.Println(i2) // 2
		fmt.Println("len", len(ch3)) // 1

		// 下記のようにデータの受信をする事もできる
		fmt.Println(<-ch3) // 3 最後に入れたデータが取り出される
		fmt.Println("len", len(ch3)) // 0

		// バッファーサイズを超えたデータを送信すると、エラーになる
		ch3 <- 1
		ch3 <- 2
		ch3 <- 3
		ch3 <- 4
		ch3 <- 5 // バッファーサイズは5なので、これでバッファーサイズいっぱい
		fmt.Println("len", len(ch3)) // 5 バッファーサイズがいっぱい
		fmt.Println(<-ch3) // 1 ここでデータを取り出すと、バッファーサイズが1空く
		fmt.Println("len", len(ch3)) // 4 バッファーサイズが1空いた
		// バッファーサイズを超えたデータを送信すると、エラー(デッドロック)になる。fatal error: all goroutines are asleep - deadlock!
		ch3 <- 6 // 上でデータを取り出したから、バッファーサイズに空きができてエラーにならない
}
