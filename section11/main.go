package main

// 構造体 埋め込み
// 構造体はフィールドに構造体を埋め込むことができる

import "fmt"

type T struct {
		// フィールド名 型
		User User // UserフィールドにUser型を埋め込む
		// 上記のように、フィールド名と型が同名になることは、Goではよくあること
}

type T2 struct {
		// また、型名を省略することもできる
		User // 型名を省略すると、フィールド名と型名は同一になる。つまり、T2はTと同じになる
}

type User struct {
		Name string
		Age int
}

func (u *User) SetName(name string) {
		u.Name = name
}

func main() {
		// Userの初期値を持ったT型の変数を定義
		var t T = T{User: User{Name: "Taro", Age: 20}}
		fmt.Println(t) // {{Taro 20}} 構造体の中に構造体が埋め込まれているので、{{}}となる

		// 構造体のフィールドにアクセスするときは、変数.フィールド名という形でアクセスする
		fmt.Println(t.User) // {Taro 20} Userフィールドの値が出力されるので、{ }となる

		// さらに、構造体に埋め込まれた構造体のフィールドにアクセスするときは、変数.フィールド名.フィールド名という形でアクセスする
		fmt.Println(t.User.Name) // Taro tに埋め込まれたUserフィールドのNameフィールドの値が出力されるので、Taroとなる

		var t2 T2 = T2{User: User{Name: "Jiro", Age: 30}}

		// 型名を省略して構造体に構造体を埋め込んだ場合、埋め込まれた構造体のフィールドにアクセスするときは、変数.埋め込まれた構造体のフィールド名という形でアクセスできる
		fmt.Println(t2.Name) // Jiro
		// 上記のように、埋め込まれた構造体のフィールド名を省略することができる
		fmt.Println(t2.User.Name) // Jiro 普通にアクセスすることもできる

		// NG例
		// fmt.Println(t.Name) // エラーになる。T型にNameフィールドがないから

		// tに埋め込まれたUserフィールドのメソッドを使うこともできる
		t.User.SetName("Saburo") // レシーバーのNameの値を変更
		fmt.Println(t.User.Name) // Saburo

		// t2のように型名を省略して構造体に構造体を埋め込んだ場合、埋め込まれた構造体のメソッドを使うときは、変数.メソッド名()という形で呼び出す
		t2.SetName("Shiro") // レシーバーのNameの値を変更
		fmt.Println(t2.Name) // Shiro
		t2.User.SetName("Goro") // レシーバーのNameの値を変更 普通にアクセスすることもできる
		fmt.Println(t2.Name) // Goro
}
