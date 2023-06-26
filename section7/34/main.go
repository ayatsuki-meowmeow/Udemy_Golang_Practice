package main

import "fmt"

// ジェネレーター
// 何らかのルールに従って、連続した値を返す
// クロージャーの応用で実装できる

func integers() func() int {
		i := 0
		return func() int {
				i++
				return i
		}
}

func main() {
		ints := integers()
		fmt.Println(ints())
		fmt.Println(ints())
		fmt.Println(ints())
		fmt.Println(ints())
		fmt.Println(ints())

		// 別のクロージャーを作成すると、iの値は初期化される
		var ints2 = integers()
		fmt.Println(ints2())
		fmt.Println(ints2())
		fmt.Println(ints2())
		fmt.Println(ints2())
		fmt.Println(ints2())
}
