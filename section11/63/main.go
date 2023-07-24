package main

// struct(構造体) メソッド
// 任意の型に特化した関数を定義するための仕組み

import "fmt"

type User struct {
		Name string
		Age int
}

// メソッドの定義
// func (レシーバー変数 レシーバーの型) メソッド名(引数) 戻り値 {
func (u User) SayName() {
		fmt.Println(u.Name) // User型のNameフィールドを出力
}

// レシーバーの型は、ポインタ型にすることもできる
// ポインタ型にすると、メソッド内でレシーバーの値を変更できる
// 参照渡しをしたいときにポインタ型にする
func (u *User) SetName(name string) {
		u.Name = name // レシーバーのNameの値を変更
}

// ポインタ型にしないと、値渡しになるので、メソッド内でレシーバーの値を変更できない
func (u User) NotSetName(name string) {
		u.Name = name // レシーバーのNameの値を変更(値渡しなので、変更されない)
}

func main() {
		var user1 User = User{Name: "Taro", Age: 20}
		// user1のアドレスを表示
		fmt.Println(&user1) // &{Taro 20}

		// メソッドを使う時は、変数.メソッド名()という形で呼び出す
		user1.SayName() // Taro

		user1.SetName("Jiro") // レシーバーのNameの値を変更
		user1.SayName() // Jiro

		user1.NotSetName("Saburo") // レシーバーのNameの値を変更(値渡しなので、変更されない)
		user1.SayName() // Jiro

		// 下記のようにはできない。エラーになる
		// &user1.SetName("Shiro")
		// user1.SayName()

		// メソッドのレシーバーは、ポインタ型にしておくのが望ましいと思われる
		// そうすると、下記のように、呼び出し側をポインタ型にする必要がない
		var user2 *User = &User{Name: "Shiro", Age: 30}
		user2.SetName("Goro") // レシーバーのNameの値を変更
		user2.SayName() // Goro

		// レシーバーが値型かポインタ型かは、メソッドで宣言されているレシーバーの型で決まる
		// よって、呼び出し側が値型かポインタ型かは関係ない
		// 下記は、ポインタ型で宣言した変数だが、メソッドで宣言されているレシーバーの型が値型なので、値渡しになる
		user2.NotSetName("Rokuro") // レシーバーのNameの値を変更(値渡しなので、変更されない)
		user2.SayName() // Goro

		// 以上より、構造体に関するメソッドのレシーバーは、ポインタ型にしておくのが望ましいと思われる
}
