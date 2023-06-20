package main

import "fmt"

// 関数

// 引数の型を指定して、返り値の型を指定する
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

// 返り値の変数名を宣言する(ここだとresult)ことで、returnの中身を省略できる
func Double(price int) (result int) {
		result = price * 2
		return
}

// 返り値のない関数
func Noreturn() {
		fmt.Println("No Return")
		return
}

func main() {
		i := Plus(1, 2)
		fmt.Println(i)

		var i2 int = Plus2(3, 4)
		fmt.Println(i2)

		var i3, i4 int = Div(4, 3)
		// fmt.Println(i3)だとエラーになる
		// 変数宣言した変数は必ず使わないといけない
		fmt.Println(i3, "余り", i4)

		// 返り値の破棄
		// 使いたくない変数は_にする事で破棄できる
		var i5, _ int = Div(4, 3)
		fmt.Println(i5) // エラーにならない

		var i6 int = Double(1000)
		fmt.Println(i6)

		Noreturn() // No Return
}
