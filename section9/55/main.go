package main

import (
	"fmt"
	"time"
)

// channel
// チャネルとゴルーチン
// チャネルは単なるキューではなく、ゴルーチン間のデータの共有のために用意されている
// 複数のゴルーチン間でチャネルを共有するプログラムが現れる

func reciever(c chan int) {
		for {
				// チャネルから整数を受け取る
				i := <- c
				// その整数を表示する
				fmt.Println(i)
		}
}

func main() {
		ch1 := make(chan int)
		var ch2 chan int = make(chan int)

		// 以下はデッドロックになる
		// fmt.Println(<-ch1)

		// 並行で走らせることで共有しつつ処理を行う
		// recieverはチャネルにデータが入るのを待っている
		go reciever(ch1)
		go reciever(ch2)

		// チャネルにデータを送る
		for i := 0; i < 100; i++ {
				// チャネルに送信すると同時に、recieverが受け取る
			  ch1 <- i
				ch2 <- i
				time.Sleep(50*time.Millisecond)
		}

		// ゴルーチンの実行順序やチャネルの受け渡しのタイミングによって、出力結果は飛び飛びになる
		// iが99の時、ch2には100が渡されるので、100が出力される
		var i int = 0
		for i < 100 {
				ch1 <- i
				// ここで待たないと、順番が保証されない
				time.Sleep(50*time.Millisecond)
				i++
				ch2 <- i
				time.Sleep(50*time.Millisecond)
		}
}
