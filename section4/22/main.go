package main

import (
	"fmt"
	"strconv"
)

// 型変換

func main() {
		// 数値型の相互変換
		var i int = 1
		fl64 := float64(i)
		fmt.Println(fl64)
		fmt.Printf("i = %T\n", i) // int
		fmt.Printf("fl64 = %T\n", fl64) // float64

		i2 := int(fl64)
		fmt.Printf("i2 = %T\n", i2) // int

		fl := 10.5
		i3 := int(fl)
		fmt.Printf("fl = %T\n", fl) // float64
		fmt.Println(fl) // 10.5
		fmt.Printf("i3 = %T\n", i3) // int
		fmt.Println(i3) // 10 四捨五入ではなく、小数点以下切り捨て

		// 文字列型から数値型への変換
		var s string = "100"
		fmt.Printf("s = %T\n", s) // string

		// strconvというパッケージのAtoiという関数は、文字列型から数値型への変換をする関数
		// 返ってくる変数は2つあるが、2つ目を _ にすることで、2つ目の変数を破棄している
		// 2つ目の変数を破棄することで、変数を必ず使わないといけないというルールを無視できる
		i4, _ := strconv.Atoi(s)
		fmt.Printf("i4 = %T\n", i4) // int
		fmt.Println(i4) // 100

		// あるいは、if文で2つ目の変数をハンドリングすることもできる
		i5, err := strconv.Atoi("A") // Aは数値に変換できない
		if err != nil {
				fmt.Println(err) // strconv.Atoi: parsing "A": invalid syntax
		}
		fmt.Println(i5) // 0

		// 数値型から文字列型への変換
		var i6 int = 200

		// strconvというパッケージのItoaという関数は、数値型から文字列型への変換をする関数
		s2 := strconv.Itoa(i6)
		fmt.Printf("s2 = %T\n", s2) // string
		fmt.Println(s2) // 200

		// 文字列からバイト配列への変換
		var h string = "Hello World"
		b := []byte(h)
		fmt.Printf("b = %T\n", b) // []uint8
		fmt.Println(b) // [72 101 108 108 111 32 87 111 114 108 100]

		// バイト配列から文字列への変換
		h2 := string(b)
		fmt.Printf("h2 = %T\n", h2) // string
		fmt.Println(h2) // Hello World
}
