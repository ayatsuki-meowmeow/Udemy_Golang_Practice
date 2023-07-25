package main

import "fmt"

// 構造体 map
// mapのキーや値に構造体を使うことができる

type User struct {
		Name string
		Age int
}

func main() {
		// mapの値に構造体を使う
		// 変数名 := map[キーの型]値の型
		m := map[int]User{
				1: {Name: "Taro", Age: 20},
				2: {"Jiro", 30},
		}
		fmt.Println(m) // map[1:{Taro 20} 2:{Jiro 30}] mapの中に構造体が入っている

		//mapのキーに構造体を使う
		// var 変数名 map[キーの型]値の型 = map[キーの型]値の型{キー: 値}
		var m2 map[User]string = map[User]string{
				{Name: "Taro", Age: 20}: "Tokyo",
				{"Jiro", 30}: "Osaka",
		}
		fmt.Println(m2) // map[{Jiro 30}:Osaka {Taro 20}:Tokyo] mapの中に構造体が入っている
		// Goではmapのキーの順序は保証されないので、キーの順序は変わる可能性がある

		// make関数を使って、mapを作成することもできる
		var m3 map[int]User = make(map[int]User)
		fmt.Println(m3) // map[]
		// m3に値を追加する
		m3[1] = User{Name: "Taro", Age: 20}
		fmt.Println(m3) // map[1:{Taro 20}]

		// for文でmapの値だけを出力する
		for _, v := range m {
				fmt.Println(v) // {Taro 20} {Jiro 30}
		}
}
