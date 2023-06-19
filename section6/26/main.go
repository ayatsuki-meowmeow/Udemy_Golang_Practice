package main

import "fmt"

// 論理演算子

func main() {
		// A && B は、AとBの両方がtrueの時にtrueになる
		fmt.Println(true && false == true) // false
		fmt.Println(true && true == true) // true
		fmt.Println(true && false == false) // true

		// A || B は、AかBのいずれかがtrueの時にtrueになる
		fmt.Println(true || false == true) // true
		fmt.Println(false || true == false) // false

		// !を使うとtrueとfalseが反転する
		fmt.Println(!true) //false
		fmt.Println(!false) //true
}
