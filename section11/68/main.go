package main

import "fmt"

// 独自型
// type は、新しい型を定義するためのキーワード
// type 型名 型

// 独自型は、あるデータの型を厳しく定義して、他のデータ型と計算等をすることを防ぐために使う

// 独自型の定義
type MyInt int // int型の独自型MyInt型を定義

// 独自型のメソッド
// func (レシーバー変数 レシーバーの型) メソッド名(引数) 戻り値 {
func (my MyInt) Print() {
		fmt.Println(my)
}

func main() {
		var mi MyInt
		fmt.Println(mi) // 0 MyInt型の初期値は0
		fmt.Printf("%T\n", mi) // main.MyInt MyInt型

		// int型とMyInt型は別の型なので、計算等をすることはできない
		// var i int = 100
		// fmt.Println(mi + i) // エラーになる

		// 独自型のメソッドを呼び出す
		mi.Print() // 0
}
