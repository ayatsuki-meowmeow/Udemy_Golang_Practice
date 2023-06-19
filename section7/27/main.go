package main

import "fmt"

// 関数

// 引数の型を指定して、返り値の型をしてする
func Plus(x int, y int) int {
		return x + y
}

// 引数の型が同じときはまとめて指定する事もできる
func Plus2(x, y int) int {
		return x + y
}

// 返り値を複数にする事もできる
func Div(x, y int) (int, int) {
		q := x / y
		var r int = x % y
		return q, r
}

func main() {
		i := Plus(1, 2)
		fmt.Println(i)

		var i2 int = Plus2(3, 4)
		fmt.Println(i2)

		var i3, i4 int = Div(4, 3)
		fmt.Println(i3, "余り", i4)
}
