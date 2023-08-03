package main

// os
// Goは何もしなければ、main関数が終了した時点で、プログラムを終了する
// しかし、os.Exit()を使うと、任意のタイミングでプログラムを強制終了することができる
// また、os.Exit()は、引数に終了ステータスを指定することができる

import (
	"fmt"
	// "log"
	"os"
)

func main() {
		// Exit
		// os.Exit(1) // exit status 1 ここでプログラムが終了する。終了ステータスは1
		// fmt.Println("Start!")

		defer func() {
				fmt.Println("defer")
		}()
		os.Exit(0) // ここでプログラムが終了する。終了ステータスは0。deferも実行されないことに注意
}
