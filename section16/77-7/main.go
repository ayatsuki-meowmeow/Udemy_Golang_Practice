package main

import (
		"fmt"
		"log"
		"os"
)

func main() {
		// ファイル操作
		// OpenFile
		// OpenFileは、ファイルを開く関数で、より詳細に指定してファイルを開くことができる
		// 第一引数には、ファイル名を指定する
		// 第二引数には、ファイルのオープンモード(フラグ)を指定する
		// 第三引数には、ファイルのパーミッションを指定する
		// OpenFileの戻り値は、ファイルとエラーを返す

		// ファイルを開くモード(フラグ)は、以下のように指定する
		// O_RDONLY: 読み込み専用
		// O_WRONLY: 書き込み専用
		// O_RDWR: 読み書き両方可能
		// O_APPEND: ファイル末尾に追記
		// O_CREATE: ファイルが存在しない場合、新規作成
		// O_TRUNC: 可能であれば、ファイルの内容をオープン時に空にする

		// ファイルを開く
		// フラグは複数指定することができる。複数指定する場合は、|でつなぐ
		// 複数指定した場合、指定した順番に処理される
		f, err := os.OpenFile("../77-5/B.txt", os.O_RDONLY|os.O_CREATE, 0666)
		if err != nil {
				log.Fatalln(err)
		}

		// 最後にファイルを閉じる
		defer f.Close()

		// ファイルの内容を受け取るためのbyte型のスライスを作成する
		bs := make([]byte, 128)
		// ファイルの内容を読み取る
		n, err := f.Read(bs)
		if err != nil {
				log.Fatalln(err)
		}
		fmt.Println(n) // 27
		fmt.Println(string(bs)) // Hello World Golang Go言語
}
