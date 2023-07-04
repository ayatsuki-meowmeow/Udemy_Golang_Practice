package main

// defer
// 関数の終了時に実行する処理を登録する

import "fmt"
import "os"

// 基本的なdeferの使い方
func TestDefer() {
		defer fmt.Println("END") // TestDefer関数の終了時に実行される
		fmt.Println("START")
}

// 複数のdeferを登録する
// deferは後に登録されたものから実行される
// 下記の場合、3, 2, 1の順で実行される
func RunDefer() {
		defer fmt.Println("1")
		defer fmt.Println("2")
		defer fmt.Println("3")
}

func main() {
		TestDefer()

		// deferで複数の処理を登録する
		// 無名関数を使って登録する
		// 一つのdeferに複数の処理をまとめたパターン
		defer func() {
				fmt.Println("------")
				fmt.Println("1")
				fmt.Println("2")
				fmt.Println("3")
		}()

		RunDefer()

		// defer文を使ったリソースの解放
		file, err := os.Create("test.txt") // osパッケージを使ってファイルを作成。osパッケージに関しては後ほど詳しく説明する
		if err != nil {
				// エラーがあればエラー内容が出力される
				fmt.Println(err)
		}
		defer file.Close() // deferを使って、fileを閉じることでリソースを解放する

		file.Write([]byte("Hello World")) // test.txtを開いてHello Worldと書き込む
}
