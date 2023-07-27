package main

import "fmt"

// interface(インターフェース)型
// カスタムエラー

// Goの組み込みのエラー型は、interface型で定義されている
// 以下のようなイメージ
// type error interface {
// 		 Error() string
// }

// カスタムエラーで独自のエラー型を定義する
type MyError struct {
		Message string
		ErrCode int
}

// カスタムエラー型にError()メソッドを実装することで、Goのソースコードにあるエラー型(インターフェース型)と同じように扱うことができる
// つまり、エラー型とカスタムエラー型が同じ性質を持つようになる
func (e *MyError) Error() string {
		return e.Message
}

// カスタムエラーを返す関数
// 返り値のerror型は、Goに組み込まれているエラー型(インターフェース型)
// しかし、実際に返されるのは、カスタムエラー型であるMyError型のポインタ
// MyError型は、Error()メソッドを実装しているので、エラー型として扱うことができる(エラー型と同じ性質を持っているため)
func RaiseError() error {
		return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
		err := RaiseError() // errはerror型だが、MyError型のポインタを返す
		fmt.Println(err.Error()) // カスタムエラーが発生しました

		// fmt.Println(err.Message) // エラー型にはMessageプロパティがないので、コンパイルエラーになる

		// 型アサーションを使って、Messageプロパティにアクセスできる
		// 型アサーションは、errがMyError型を持っているかどうかをチェックして、MyError型ならeにerrを代入し、okにtrueを代入する
		// errがMyError型を持っていない場合は、eにnilを代入し、okにfalseを代入する
		e, ok := err.(*MyError) // eはMyError型、okはbool型、errはerror型
		if ok {
				fmt.Println(e.Message) // カスタムエラーが発生しました
				fmt.Println(e.ErrCode) // 1234
		} else {
				fmt.Println("カスタムエラーではありません")
		}
}
