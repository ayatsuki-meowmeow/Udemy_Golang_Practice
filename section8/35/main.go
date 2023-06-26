package main

// if文
// 条件分岐

import "fmt"

func main() {
		var a int = 1
		if a == 2 {
				// aが2の場合trueになるので、twoが出力される
				fmt.Println("two")
		} else if a == 1 {
				// aが1の場合trueになるので、oneが出力される
				fmt.Println("one")
		} else {
				// aが与えた条件に当てはまらない場合falseになるので、I don't knowが出力される
				fmt.Println("I don't know")
		}

		// 簡易文付きif文
		// if文の前に簡易文を書くことができる
		// varでの変数宣言はできない
		if b := 100; b == 100 {
				fmt.Println("one hundred")
		}

		// 簡易文付きif文のスコープ
		// 簡易文付きif文のスコープはif文の中だけ
		var x int = 0
		if x := 2; true {
				fmt.Println(x) // 2
		}
		fmt.Println(x) // 0
}
