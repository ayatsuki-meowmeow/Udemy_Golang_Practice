package main

import (
		"fmt"
		"os"
)

func main() {
		// Args
		// ビルドして、実行ファイルを作成する
		// go build -o 実行ファイル名 ファイル名
		// 今回だと、go build -o getargs main.go

		// 実行ファイルを実行する
		// ./実行ファイル名 引数1 引数2 引数3 ... のように実行する(この引数をコマンドライン引数という)
		// 今回だと、./getargs 123 456 789 のような感じ

		// Argsは、string型のスライスで、任意のコマンドライン引数が格納されている
		// os.Args[0]には、実行ファイル名が格納されている
		fmt.Println(os.Args[0]) // ./getargs インデックス番号0の要素には、実行ファイル名が格納されている
		fmt.Println(os.Args[1]) // インデックス番号1の要素には、コマンドライン引数が格納されている
		fmt.Println(os.Args[2])
		fmt.Println(os.Args[3])

		// os.Argsの要素数を表示
		fmt.Printf("length=%d\n", len(os.Args))
		// os.Argsの要素を全て表示
		for i, v := range os.Args {
				fmt.Println(i, v) // インデックス番号と、要素の値が表示される
		}
}
