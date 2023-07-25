package main

import "fmt"

// 構造体 型コンストラクタ
// Goではコンストラクタという概念はないが、コンストラクター関数を使うことがよくある

type User struct {
		Name string
		Age int
}

// コンストラクター関数
// New型名という関数名にするのが一般的
// 戻り値は、*型名にするのが一般的(ポインタ型にするのが一般的)
// 引数は、構造体のフィールドになるようにするのが一般的
// func New型名(引数) *型名 {
// 		return &型名{引数} // &型名{引数}という形で構造体のポインタを返す
// }
func NewUser(name string, age int) *User {
		return &User{Name: name, Age: age}
}

func main() {
		user1 := NewUser("Taro", 20)
		// user1 := &User{Name: "Taro", Age: 20} と同じ
		fmt.Println(user1) // &{Taro 20}
		fmt.Println(*user1) // {Taro 20} ポインタ型なので、デリファレンスすると値が出力される(user1の実体の値が出力される)

		var user2 *User = NewUser("Jiro", 30)
		// user2 := &User{Name: "Jiro", Age: 30} と同じ
		fmt.Println(user2) // &{Jiro 30}
}
