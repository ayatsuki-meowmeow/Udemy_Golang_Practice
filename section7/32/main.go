package main

import "fmt"

// クロージャー

// 返り値はstring型の引数を取りstring型の値を返す関数
func Later() func(string) string {
		var store string // storeを空文字で定義する
		// string型のfunc(next string)を返す
		return func(next string) string {
				s := store // sにstoreを代入する
				store = next // storeにnext(引数)を代入する
				return s // sを返す(この段階ではsの値はnextを代入する前のstore)
		}
}

func main() {
		f := Later()
		fmt.Println(f("Hello"))
		fmt.Println(f("My"))
		fmt.Println(f("name"))
		fmt.Println(f("is"))
		fmt.Println(f("Golang"))
}

// Later 関数を呼び出すと、その返り値として新しい関数（クロージャ）が得られる
// Later を実行するタイミングで、storeは初期化されない
