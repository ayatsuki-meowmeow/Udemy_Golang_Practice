package main

import "fmt"

// byte(uint8)型

func main() {
		byteA := []byte{72, 73}
		// []←スライス、配列を表現できる。詳しくはスライスの講義で
		fmt.Println(byteA)

		// byte型を文字列に変換する
		fmt.Println(string(byteA))

		// 文字列をbyteのスライスに変換する
		c := []byte("HI")
		fmt.Println(c)

		// スライスを文字列に変換する
		fmt.Println(string(c))
}
