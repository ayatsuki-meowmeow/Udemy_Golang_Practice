package main

// for
// 繰り返し処理

import "fmt"

func main() {
		// 無限ループ
		// ctrl + c で強制終了
		// for {
		// 		fmt.Println("Loop")
		// }

		var i int = 0
		for {
				i++
				if i == 3 {
						// iが3になったらループを抜ける
						// breakはfor文を強制終了して抜ける
						break
				}
				fmt.Println("Loop")
				// Loopが2回出力される
		}

		// 条件付きfor文
		var point int = 0
		// pointが10未満の場合ループを繰り返す
		// 条件がtrueだったらループを繰り返す
		// 条件がfalseになったらループを抜ける
		for point < 10 {
				fmt.Println(point)
				point++
		}
}
