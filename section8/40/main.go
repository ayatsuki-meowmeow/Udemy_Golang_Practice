package main

// ラベル付きfor文
// ラベルを使って、任意の位置まで処理を飛ばすことができる

import "fmt"

func main() {
		// 下記のコードだと無限ループになる
		// for {
		// 		for {
		// 				for {
		// 						fmt.Println("START")
		// 						break 一番近くのfor文を抜ける
		// 				}
		// 				fmt.Println("処理をしない")
		// 		}
		// 		fmt.Println("処理をしない")
		// }
		// fmt.Println("END")

		// ラベル付きfor文を使うと、一気に抜けることができる
Loop:
		for { // このfor文にLoopというラベルが付いている
				for {
						for {
								fmt.Println("START")
								break Loop // Loopに書かれたfor文を抜ける
						}
						fmt.Println("処理をしない") // この処理は実行されない
				}
				fmt.Println("処理をしない") // この処理は実行されない
		}
		fmt.Println("END") // Loopに書かれたfor文を抜けたら、この処理が実行される

		// 通常のfor文。全てを出力する
		for i := 0; i < 3; i++ {
				for j := 1; j < 3; j++ {
					fmt.Println(i, j, i*j)
				}
		}

		fmt.Println("----")

		fmt.Println("失敗例")
		for i := 0; i < 3; i++ {
				for j := 1; j < 3; j++ {
						if j > 1 {
								// jが1より大きい場合、処理をスキップする
								// continueはfor文の処理をスキップする
								// continueで抜けるのは一番近いfor文
								continue
						}
						fmt.Println(i, j, i*j)
				}
				// continueで抜けた後、以下の処理がされてしまう
				fmt.Println("処理をしない")
		}

		fmt.Println("成功例")
Loop2:
		for i := 0; i < 3; i++ { // このfor文にLoop2というラベルが付いている
				for j := 1; j < 3; j++ {
						if j > 1 {
								// jが1より大きい場合、処理をスキップする
								// continueでLoop2に書かれたfor文を抜ける
								// for i := 0; i < 3; i++ の行まで戻る
								continue Loop2
						}
						fmt.Println(i, j, i*j)
				}
				// continueで抜けた後、以下の処理がされない
				fmt.Println("処理をしない")
		}
}
