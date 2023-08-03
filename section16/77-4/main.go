package main

import (
		"log"
		"os"
)

func main() {
		// ファイル操作
		// Open
		f, err := os.Open("A.txt") // os.Open()は、ファイルを開く関数
		if err != nil { // A.txtが存在しない場合、エラーが発生する
				log.Fatalln(err) // open A.txt: no such file or directory exit status 1
		}

		defer f.Close() // 関数が終了する時に、ファイルを閉じる
}
