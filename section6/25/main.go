package main

import "fmt"

// 比較演算子

func main() {
		// A == B でAとBが等しいかを比較する
		fmt.Println(1 == 1) // true

		fmt.Println(1 == 2) // false

		// A <= B でBがA以上かどうか比較する
		fmt.Println(4 <= 8) // true

		// A >= B でBがA以下かどうか比較する
		fmt.Println(1 >= 8) // false

		fmt.Println(1 < 8) // true
		fmt.Println(3 > 1) // true

		// 論理値をそのまま比較する
		fmt.Println(true == false) // false

		// A != B でAとBが等しくないかを比較する
		fmt.Println(true != false) // true

		// fmt.Println("1" == 1) 型が違うとエラーになる
}
