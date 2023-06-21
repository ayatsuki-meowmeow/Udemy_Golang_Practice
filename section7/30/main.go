package main

import "fmt"

// 関数を引数に取る関数

// 引数であるfというfunc()を実行する関数
func CallFunction(f func()) {
		f()
}

func main() {
		// CallFunctionに引数として無名関数を渡す
		CallFunction(func() {
				fmt.Println("I am a function")
		})
}
