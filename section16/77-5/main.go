package main

import (
		"os"
)

func main() {
		// ファイル操作
		// Create
		// Create関数は、ファイルを新規作成する

		f, _ := os.Create("B.txt") // B.txtを新規作成する。ファイルが存在する場合、ファイルを削除して、再度新規作成する。
		// 上記だと、Create関数で作成したファイルをfで受け取っている

		// f(作成したファイル)に対して、Writeメソッドを使って、ファイルに書き込む
		// Writeメソッドを使って、byte型のスライスを書き込む
		f.Write([]byte("Hello World\n"))

		// WriteAtメソッドを使って、ファイルの指定した位置に書き込む
		// ここでは、12バイト目に書き込む
		f.WriteAt([]byte("Golang\n"), 12)

		// ファイルの末尾にオフセットを移動する
		f.Seek(0, os.SEEK_END)

		// WriteStringメソッドを使って、文字列を書き込む
		// 上記でファイルの末尾に移動しているので、ファイルの末尾に書き込まれる
		f.WriteString("Go言語")
}
