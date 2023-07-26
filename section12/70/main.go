package main

import "fmt"

// interface(インターフェース)型
// 最もポピュラーな使い方。異なる型に共通の性質を付与する
// インターフェース（interface）を使って異なる型に共通の性質を付与し、それらをsliceにまとめて同じように処理する

type Person struct {
		Name string
		Age int
}

func (p *Person) ToString() string {
		// Sprintfは、フォーマットを指定して、文字列を返す関数
		// %vに、p.Nameとp.Ageを入れる
		return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
		Number string
		Model string
}

func (c *Car) ToString() string {
		return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

// Person型とCar型の両方をsliceに入れて、ToString()を呼び出したい
// しかし、Person型とCar型は、異なる型なので、sliceに入れることができない
// そこで、interface型を使う
type Stringfy interface {
		// ToString()メソッドをStringfy型にまとめる
		ToString() string
}

func main() {
		var vs []Stringfy = []Stringfy{
				&Person{Name: "Taro", Age: 20},
				&Car{Number: "111-222", Model: "AB-1234"},
		}
		// 通常、Person型とCar型は、異なる型なので、sliceに入れることができない
		// しかし、共通のToString()メソッドを通じて、Stringfy型にまとめることができる

		for _, v := range vs {
				// ここで、vはStringfy型なので、v.ToString()を呼び出すことができる
				fmt.Println(v.ToString()) // Name=Taro, Age=20 Number=111-222, Model=AB-1234
		}

		// このように、型の性質を抽出したインターフェースを定義することで、Goの厳密な型のシステムに柔軟性を与えることができる。

		for i:=0; i < len(vs); i++ {
				// ここで、vs[i]はStringfy型なので、vs[i].ToString()を呼び出すことができる
				fmt.Println(vs[i].ToString()) // Name=Taro, Age=20 Number=111-222, Model=AB-1234
		}
}

// Go言語では、あるインターフェースを実装するために必要なのは、そのインターフェースで定義されているすべてのメソッドを対象の型で実装すること
// 今回は、Stringfy型には、ToString()メソッドが定義されているので、Person型とCar型は、ToString()メソッドを実装する必要がある
