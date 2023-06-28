package main

import "fmt"

// switch
// 式スイッチ
// 条件分岐の一つ
// 式を評価して処理を分岐する

func main() {
		var n int = 1
		// var n string = "A" 型が違うからエラーになる
		switch n {
		case 1, 2:
				// nが1または2の場合
				fmt.Println("1 or 2")
		case 3, 4:
				// nが3または4の場合
				fmt.Println("3 or 4")
		default:
				// どれにも当てはまらない場合
				// defaultはひとつだけ書ける
				fmt.Println("I don't know")
		}

		// 下記のように書く事もできる
		// この場合、varは使えない。 syntax error: var declaration not allowed in switch initializer
		// このように書くと、変数nはswitch文の中でしか参照できない
		switch n := 2; n {
		case 1, 2:
				fmt.Println("1 or 2")
		case 3, 4:
				fmt.Println("3 or 4")
		default:
				fmt.Println("I don't know")
		}

		// 下記のように書く場合は、switchに与える変数を省略できる
		var n2 int = 6
		switch {
		case n2 > 0 && n2 < 4:
				fmt.Println("0 < n2 < 4")
		case n2 > 3 && n2 < 7:
				fmt.Println("3 < n2 < 7")
		}

		// 下記のように、列挙型の式と演算子を使った式が混在しているとエラーになる
		// switch n := 2; n {
		// case 1, 2:
		// 		fmt.Println("1 or 2")
		// case 3, 4:
		// 		fmt.Println("3 or 4")
		// case n > 0 && n < 4:  cannot convert n > 0 && n < 4 (untyped bool value) to int
		// 		fmt.Println("0 < n < 4")
		// default:
		// 		fmt.Println("I don't know")
		// }

		// 下記のように、すればエラーにならない
		// 複数のcaseでtrueになる場合、最初にtrueになったcaseの処理が実行される
		switch n := 2;{
		case n == 1 || n == 2:
				fmt.Println("1 or 2")
		case n == 3 || n == 4:
				fmt.Println("3 or 4")
		case n > 0 && n < 8:
				fmt.Println("0 < n < 8")
		default:
				fmt.Println("I don't know")
		}
}
