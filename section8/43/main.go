package main

// go goroutine
// goで並行処理を行うための仕組み
// 並行処理を簡単に実装できる

import "time"

func sub() {
		for {
				println("sub loop")
				time.Sleep(100 * time.Millisecond)
		}
}

func main() {
		go sub()
		go sub()

		// goがない状態だと、sub()が終了するまでmain()の処理が実行されない→永遠に実行されない
		// goがある状態だと、sub()とmain()が並行して実行される
		for {
				println("main loop")
				time.Sleep(200 * time.Millisecond)
		}
}
