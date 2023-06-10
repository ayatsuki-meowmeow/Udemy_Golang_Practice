package main

import "fmt"

// interfaceとnil
// interface型はあらゆる型と互換性がある特殊な型。初期値はnil
// nilは何の値も持っていない

func main() {
		var x interface{}
		fmt.Println(x) // <nil>

		// xにはあらゆる型の値を代入できる
		x = 1
		fmt.Println(x) // 1

		x = 3.14
		fmt.Println(x)

		x = "A"
		fmt.Println(x)

		x = [3]int{1, 2, 3}
		fmt.Println(x)

		// interface型では、データ特有の計算や演算はできない
		x = 2
		// fmt.Println(x + 3)
		// invalid operation: x + 3 (mismatched types interface{} and int)
		// interface型とint型では計算できない
}

// interface型の詳細については、後の講義で行う
