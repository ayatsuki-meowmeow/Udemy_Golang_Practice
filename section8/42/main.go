package main

import "fmt"

// panic recover
// 例外処理に似た仕組み
// ランタイムを強制的に停止させる機能を持つので、エラーハンドリングをした方が良い
// プログラムを強制的に終了させるので、使うのは避けるべき(かも)
// recoverを使うことで、panicが発生してもプログラムを終了させずに処理を続行させることができる
// recoverは原則としてdeferと組み合わせて使う

func main() {
		defer func() {
				if x := recover(); x != nil { // recover()はpanic()が発生したときにpanic()の引数を取得する
						fmt.Println(x) // runtime error
				}
		}()
		panic("runtime error") // panicは実行中のプログラムを強制終了させる
		fmt.Println("Start") // panicが発生したので、この処理は実行されない
}
