package main

import "fmt"

// 算術演算子

func main() {
		fmt.Println(1 + 3)
		fmt.Println("ABC" + "DE") // ABCDE 文字列の結合に使える

		fmt.Println(5 - 4)

		fmt.Println(5 * 4)

		fmt.Println(60 / 3)

		fmt.Println(9 % 4) // 9/4の余りである1が出力される

		n := 0
		n += 4 // n = n + 4 と同じ
		fmt.Println(n)
		n++ // n = n + 1 と同じ
		fmt.Println(n)
		n-- // n = n - 1 と同じ
		fmt.Println(n)

		var s string = "ABC"
		s += "DEF" // "ABC" + "DEF" と同じ
		fmt.Println(s)
}
