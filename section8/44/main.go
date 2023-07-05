package main

import "fmt"

// init
// init関数はmain関数より先に実行される
// 初期化処理を確実に行うためにinit関数を使う
// 複数のinit関数がある場合、上から順に実行される
// init関数を複数定義することはで定義することはできるが、あまり意味はないのでまとめて定義すれば良い

func init() {
	fmt.Println("Init1")
}

func main() {
	fmt.Println("Main")
}

func init() {
	fmt.Println("Init2")
}

// Init1
// Init2
// Main
