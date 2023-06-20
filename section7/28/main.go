package main

import "fmt"

// 無名関数

func main() {
		// 関数名がないだけで、関数とほとんど同じ
		var f = func(x, y int) int {
				return x + y
		}
		var i int = f(1, 2)
		fmt.Println(i)

		f2 := func(x, y int) int {
				return x + y
		}
		i2 := f2(1, 2)
		fmt.Println(i2)

		// 無名関数を宣言してそのまま引数を渡す事もできる
		var i3 = func(x, y int) int {
				return x + y
		}(1, 2) // (1, 2)が引数
		fmt.Println(i3)
}
