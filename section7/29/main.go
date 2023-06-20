package main

import "fmt"

// 関数を返す関数

// func()が返り値の型に該当する
// func()型だと考えるとわかりやすい
func ReturnFunc() func() {
		return func() {
				fmt.Println("I am a function")
		}
}

func main() {
		var f = ReturnFunc()
		f()
}

// fには
// func() {
// 		fmt.Println("I am a function")
// }
// という無名関数が入っていることになる
