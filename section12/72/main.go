package main

import "fmt"

// interface(インターフェース)型
// fmt.Stringer
// fmtパッケージに定義されているStringer型は、interface型で定義されている
// Stringer型は文字列を返すString()メソッドのみを持つ
// 以下のようなイメージ
// type Stringer interface {
// 		String() string
// }

// String()メソッドを持つ型を定義することで、Stringer型として扱うことができる
type Point struct {
		A int
		B string
}

// String()メソッドを実装する
// func (レシーバー変数 レシーバーの型) メソッド名(引数) 戻り値 {
func (p *Point) String() string {
		// Sprintfは、フォーマットを指定して、文字列を返す関数
		// %vに、p.Aとp.Bを入れる
		return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
		var p *Point = &Point{1, "test"}
		// ↑は、p := &Point{1, "test"}と同じ

		// String()メソッドを実装しているので、Stringer型として扱うことができる

		fmt.Println(p.String()) // <<1, test>>
		fmt.Println(p) // <<1, test>>
		// Go言語では、fmtパッケージのPrintln()関数は、引数がStringer型の場合、String()メソッドを呼び出して、文字列を取得してから出力する
		// そのため、pをそのまま渡しても、p.String()と同じ結果になる

		fmt.Println(*p) // {1 test}
		// *pは、Point型なので、Stringer型ではない
		// そのため、*pを渡すと、String()メソッドは呼び出されず、そのまま出力される
}
