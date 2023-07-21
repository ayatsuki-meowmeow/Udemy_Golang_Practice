package main

import "fmt"

// 構造体
// オブジェクト指向のクラスのようなもの
// 複数の任意の型の値を一つにまとめたもの

// 構造体の定義
// type 構造体名 struct {
type User struct {
		// フィールド名 型
		Name string // Nameフィールド string型
		Age int // Ageフィールド int型
		// フィールド名は大文字で始める必要がある

		// フィールドをまとめて定義することもできる
		// X, Y int // X, Yフィールド int型
}

// 引数として、構造体を受け取る関数
// 値渡しになる
func UpdateUser(user User) {
		user.Name = "A"
		user.Age = 1000
}

// 引数として、構造体のポインタを受け取る関数
// 参照渡しになる
func UpdateUserV2(user *User) {
		user.Name = "B"
		user.Age = 2000
}

func main() {
		// user型を使って変数を定義
		var user1 User // user1変数はUser型として定義されている
		fmt.Println(user1) // { 0} フィールドの初期値が出力される
		// 構造体は{}で表示される。フィールドの値が{}内に表示される

		// 構造体のフィールドを更新
		user1.Name = "Taro"
		user1.Age = 20
		fmt.Println(user1) // {Taro 20}

		// 暗黙的な定義
		user2 := User{} // user2変数はUser型として定義されている
		fmt.Println(user2) // { 0}
		user2.Name = "Jiro"
		fmt.Println(user2) // {Jiro 0}

		// 暗黙的な定義のときにフィールドの値を指定することもできる
		user3 := User{Name: "Saburo", Age: 30} // 初期値を持っている
		fmt.Println(user3) // {Saburo 30}

		// 明示的な定義でもフィールドの値を指定することができる
		var user4 User = User{Name: "Shiro", Age: 40}
		fmt.Println(user4) // {Shiro 40}

		// フィールドを指定しないで、フィールドの値を指定することもできる
		// フィールドを指定しないときは、構造体のフィールドの順番に値を指定する
		user5 := User{"Goro", 50}
		var user6 User = User{"Rokuro", 60}
		fmt.Println(user5) // {Goro 50}
		fmt.Println(user6) // {Rokuro 60}

		// NG例
		// user7 := User{70, "Shichiro"}
		// fmt.Println(user7) // エラーになる

		// Nameフィールドだけ、初期値を指定することもできる
		user8 := User{Name: "Hachiro"}
		fmt.Println(user8) // {Hachiro 0}
		// 下記のような書き方はできない
		// user9 := User{"Kichiro"}
		// fmt.Println(user9) // エラーになる

		// new関数を使って構造体を生成することもできる
		// new関数は、構造体のポインタを返す
		user9 := new(User)
		fmt.Println(user9) // &{ 0}
		fmt.Println(*user9) // { 0}
		// varを使って宣言するなら、アドレス演算子をつけてポインタ型のユーザー型で宣言する
		var user10 *User = new(User)
		fmt.Println(user10) // &{ 0}
		fmt.Println(*user10) // { 0}

		// アドレス演算子を使って宣言するなら、new関数を使わなくてもポインタ型のユーザー型で宣言することができる
		var user11 *User = &User{}
		fmt.Println(user11) // &{ 0}
		fmt.Println(*user11) // { 0}
		user12 := &User{}
		fmt.Println(user12) // &{ 0}
		fmt.Println(*user12) // { 0}

		// 関数の引数として構造体を使うときに、ポインタ型で宣言することが多い
		UpdateUser(user1)
		fmt.Println(user1) // {Taro 20} UpdateUser関数でuser1の値が変更されない。値渡しになるため
		// fmt.Println(*user1) // エラーになる
		UpdateUserV2(user9)
		fmt.Println(user9) // &{B 2000} UpdateUserV2関数でuser9の値が変更される。参照渡しになるため
		fmt.Println(*user9) // {B 2000}

		// user1であっても、アドレス演算子をつけてポインタ型で関数の引数として渡すと、参照渡しになる
		UpdateUserV2(&user1)
		fmt.Println(&user1) // &{B 2000} UpdateUserV2関数でuser1の値が変更される。参照渡しになるため
		fmt.Println(user1) // {B 2000}
}
