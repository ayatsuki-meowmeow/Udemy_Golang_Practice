package main

import (
		"log"
		"os"
)

func main() {
		//log.Fatal()
		// log.Fatal()は、os.Exit(1)と同じように、プログラムを終了する
		_, err := os.Open("A.txt") // os.Open()は、ファイルを開く関数
		if err != nil {
				// log.Fatal()は、何らかのエラーが発生した場合に、プログラムを終了する
				// 引数に与えられた値を出力した後、os.Exit(1)を実行して、プログラムを終了する
				log.Fatalln(err) // open A.txt: no such file or directory exit status 1
		}
}
