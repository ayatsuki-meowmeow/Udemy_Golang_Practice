package main

import (
	  "fmt"
		"log"
		"os"
)

func main() {
		// ファイル操作
		// Read
		f, err := os.Open("B.txt") // os.Open()は、ファイルを開く関数。fは、ファイルを受け取る変数
		if err != nil { // B.txtが同じディレクトリーに存在しない場合、エラーが発生する
				log.Fatalln(err) // open B.txt: no such file or directory exit status 1
		}

		defer f.Close() // 関数が終了する時に、ファイルを閉じる

		bs := make([]byte, 128) // ファイルの内容を受け取るためのbyte型のスライスを作成する

		// Readメソッドを使って、ファイルの内容を読み込む
		// byteのスライスに、ファイルの内容が入る
		// n(戻り値)には、読み込んだバイト数が入る
		// errには、エラーが入る
		n, err := f.Read(bs)
		if err != nil {
				log.Fatalln(err)
		}
		fmt.Println(n)
		fmt.Println(string(bs)) // string型に変換して出力する

		bs2 := make([]byte, 128) // ファイルの内容を受け取るためのbyte型のスライスを作成する

		// ReadAtメソッドを使って、ファイルの指定した位置から内容を読み込む
		// ReadAtメソッドの第一引数には、ファイルの内容を受け取るbyte型のスライスを指定する
		// ReadAtメソッドの第二引数には、ファイルの読み込み開始位置(オフセット)を指定する
		n2, err := f.ReadAt(bs2, 6)
		// すでにReadメソッドでファイルの内容を読み込んでいるので、ReadAtメソッドで読み込もうとすると、
		// すでにファイルの最後に到達しているため、EOF(End Of File)エラーが発生する
		// したがって、errには、EOF(End Of File)エラーが入る
		// 今回は、errを無視してn2とbs2を出力する
		// if err != nil {
		// 		log.Fatalln(err) // EOF
		// }
		fmt.Println(n2)
		fmt.Println(string(bs2)) // string型に変換して出力する
}
