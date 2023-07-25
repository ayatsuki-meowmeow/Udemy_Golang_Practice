package main

import "fmt"

// 構造体 スライス
// 構造体のスライスはよく使われる

type User struct {
		Name string
		Age int
}

// User型をスライスとして扱う
// type スライス名 []型名
type Users []*User // Users型は、User型のポインタのスライス

// 下記のUsers2型はUsers型と同じだが、上のUsersの定義の仕方の方が望ましい
type Users2 struct {
		Users2 []*User // Users型は、User型のポインタのスライス
}

func main() {
		user1 := User{Name: "Taro", Age: 20}
		user2 := User{Name: "Jiro", Age: 30}
		user3 := User{Name: "Saburo", Age: 40}
		user4 := User{Name: "Shiro", Age: 50}

		// 複数のUser型の値をUsers型に格納する
		// まず、Users型の変数を定義
		users := Users{}
		// 次に、appendを使ってUsers型の変数にUser型の値を格納する
		// 変数名 = append(変数名, 値)
		users = append(users, &user1) // Users型はUser型のポインタのスライスなので、&をつける
		users = append(users, &user2)
		users = append(users, &user3, &user4) // appendは複数の値を追加することもできる(可変長引数)

		// Users型の変数の中身を出力
		fmt.Println(users) // [0xc0000b4000 0xc0000b4020 0xc0000b4040 0xc0000b4060] ポインタのアドレスが出力される

		// Users型の変数の中身をforで出力
		for _, user := range users {
				fmt.Println(user) // &{Taro 20} &{Jiro 30} &{Saburo 40} &{Shiro 50}
				fmt.Println(*user) // {Taro 20} {Jiro 30} {Saburo 40} {Shiro 50}
		}

		// make関数を使って、構造体のスライスを生成することもできる
		users2 := make([]*User, 0) // Users型の変数を定義
		users2 = append(users2, &user1, &user2) // Users型はUser型のポインタのスライスなので、&をつける

		// Users型の変数の中身を出力
		for _, user := range users2 {
				fmt.Println(user) // &{Taro 20} &{Jiro 30}
				fmt.Println(*user) // {Taro 20} {Jiro 30}
		}

		// そもそもUser型の要素を宣言するときに、ポインタ型にしておけば、appendのときに&をつける必要がない
		user5 := &User{Name: "Goro", Age: 60}
		user6 := &User{Name: "Rokuro", Age: 70}

		users3 := []*User{user5} // Users型の変数を定義。初期期にuser5を格納
		for _, user := range users3 {
				fmt.Println(user) // &{Goro 60}
				fmt.Println(*user) // {Goro 60}
		}
		users3 = append(users3, user6) // appendでuser6を追加
		for _, user := range users3 {
				fmt.Println(user) // &{Goro 60} &{Rokuro 70}
				fmt.Println(*user) // {Goro 60} {Rokuro 70}
		}

		// 古典的for文でも同じことができる
		for i := 0; i < len(users3); i++ {
				fmt.Println(users3[i]) // &{Goro 60} &{Rokuro 70}
				fmt.Println(*users3[i]) // {Goro 60} {Rokuro 70}
		}
}
