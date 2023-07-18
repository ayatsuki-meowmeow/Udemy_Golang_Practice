package main

// channel
// for
// チャネルから受信し続けるという処理で有効に使える
// チャネルでfor文を使う時は、値を送信し終わったら、チャネルをcloseするのが基本

func main() {
		var ch1 chan int = make(chan int, 3)
		ch1 <-1
		ch1 <-2
		ch1 <-3
		// チャネルをcloseする
		close(ch1)

		// for文でchannelからデータを受信して出力する
		// closeしていなかったら、channelが空になった時に、デッドロックになる
		for i:= range ch1 {
				println(i)
		}
}
