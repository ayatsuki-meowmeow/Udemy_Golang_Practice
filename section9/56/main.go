package main

import "fmt"
import "time"

// channel
// close
// make関数で生成したチャネルは、最初はopen状態
// チャネルをcloseすると、チャネルはclose状態になる

func reciever(name string, ch chan int) {
		for {
				i, ok := <-ch
				// チャネルがcloseされるとokにfalseが入る
				if !ok { // if ok == false と同じ
						break // チャネルがcloseされたらループを抜ける
				}
				fmt.Println(name, i)
		}
		// 終了したことを示すために、終了したことを表示する
		fmt.Println(name + " is done.")
}

func main() {
		var ch1 chan int = make(chan int, 5)

		ch1 <- 1
		ch1 <- 2

		// チャネルをcloseする
		// close(ch1)

		// closeしたチャネルにデータを送信すると、panicが発生する
		// ch1 <- 1

		// closeしたチャネルからデータを受信すると、データが受信できる
		fmt.Println(<-ch1) // 1

		// 受信は2つの変数を割り当てることができる仕組みを利用している
		// 1つ目の変数は受け取った値
		// 2つ目の変数はチャネルがopenされているかどうかの真偽値。openならtrue、closeならfalse
		// 厳密には、2つ目の変数はチャネルのバッファー内にデータがなくて、かつ、チャネルがcloseされている場合にfalseになる
		i, ok := <-ch1
		fmt.Println(i, ok) // 2 true

		// i2, ok := <-ch1
		// fmt.Println(i2, ok) // 0 false

		// gorutineとclose

		// どのrecieverが処理するかは実行する度に変わる
		// したがって、出力結果も毎回変わる
		go reciever("1.goroutine", ch1)
		go reciever("2.goroutine", ch1)
		go reciever("3.goroutine", ch1)

		for i := 0; i < 100; i++ {
				ch1 <- i
		}
		// forループが終わった後に、チャネルをcloseする
		close(ch1)
		// gorutineが終了するのを待つために、3秒待つ
		// 本来はsyncパッケージを使う
		time.Sleep(3 * time.Second)
}
