package main

import "fmt"

// channel
// select

func main() {
		// 複数のチャネルを動かすときに、いずれかのチャネルが受信できなくなると他のプログラムも停止してしまう
		// 失敗例
		// ch1 := make(chan int, 2)
		// var ch2 chan string = make(chan string, 2)

		// ch2 <- "A" // ch2に文字列を送る

		// var v1 int = <- ch1 // ch1は空だからv1には何も入らず、デッドロックになる
		// var v2 string = <- ch2 // ここに辿り着けない
		// fmt.Println(v1)
		// fmt.Println(v2)

		// 解決するためにselect文を使って、二つのチャネルの処理を分岐させる
		ch1 := make(chan int, 2)
		var ch2 chan string = make(chan string, 2)

		ch2 <- "A" // ch2に文字列を送る
		ch1 <- 1 // ch1に整数を送る
		ch2 <- "B" // ch2に文字列を送る
		ch1 <- 2 // ch1に整数を送る

		// select文を使って、ch1とch2のどちらから受信できるかによって処理を分岐させる
		// caseには、チャネルに対する処理を書く
		// switch文とは違って最初に成立したcaseの処理だけが実行されるわけではなく、どの処理が実行されるかはランダム
		// ランダムにしないと、片方のチャネルに偏ってしまう
		select {
		case v1 := <- ch1: // ch1から受信できる場合
				fmt.Println(v1 + 1000)
		case v2 := <- ch2: // ch2から受信できる場合
				fmt.Println(v2 + "!?")
		default: // どちらも受信できない場合
				fmt.Println("どちらも受信できません")
		}
}
