package main

// switch
// 型スイッチ
// interface型の活用

import "fmt"

func anything(a interface{}) {
		// interface{}型はどんな型でも入れられるが、数値の計算や文字列の結合などはできない
		// 型が失われているから
		fmt.Println(a)
}

func anything2(a interface{}) {
		// 型アサーションを使って型を復元する
		// 型を復元しているので、数値の計算や文字列の結合などができる
		switch v := a.(type) { // vにanyの値を代入する
		case string:
				fmt.Println(v + "!?")
		case int:
				fmt.Println(v + 10000)
		}
}

// 型アサーションを使って型を復元すれば、数値の計算や文字列の結合などができる

func main() {
		anything("aaa")
		anything(1)

		var x interface{} = 3
		// 型アサーションの実行
		i := x.(int) // xの中身をint型に変換してiに代入する
		var i2 int = x.(int) // 上記と同じ
		fmt.Println(i + 2) // 5
		fmt.Println(i2 + 4) // 7
		// fmt.Println(x + 2)  xはinterface{}型なので、数値の計算はできない invalid operation: x + 2 (mismatched types interface{} and int)

		// var f float64 = x.(float64)  xはint型のデータなのでfloat64では復元できずランタイムエラーになる panic: interface conversion: interface {} is int, not float64
		// fmt.Println(f)

		f, isFloat64 := x.(float64) // xをfloat64型に変換してfに代入する、変換できたかどうかをisFloat64に代入する
		fmt.Println(f, isFloat64) // 0 false(変換に失敗した)

		i3, isInt := x.(int) // xをint型に変換してi3に代入する、変換できたかどうかをisIntに代入する
		fmt.Println(i3, isInt) // 3 true(変換に成功した)

		// 2つの返り値を使った型アサーションと条件分岐を使うことで、インターフェース型で引数を渡して様々な型に対応した処理を分岐して実行することができる
		if x == nil {
				fmt.Println("None")
		} else if i, isInt := x.(int); isInt { // isIntがtrueの場合に処理を実行する
				fmt.Println(i, "x is Int")
		} else if s, isString := x.(string); isString { // isStringがtrueの場合に処理を実行する
				fmt.Println(s, "x is String")
		} else {
				fmt.Println("I don't know")
		}

		// switch文で型アサーションを使う
		// if文より簡単に書ける
		switch x.(type) {
		case int: // xの型がint型の場合
				fmt.Println("int")
		case string: // xの型がstring型の場合
				fmt.Println("string")
		default: // どれにも当てはまらない場合
				fmt.Println("I don't know")
		}

		// switch文で値も取得する
		switch v := x.(type) { // vにxの値を代入する
		case bool:
				fmt.Println(v, "bool")
		case int:
				fmt.Println(v, "int")
		case string:
				fmt.Println(v, "string")
		}

		anything2("aaa")
		anything2(1)
}
