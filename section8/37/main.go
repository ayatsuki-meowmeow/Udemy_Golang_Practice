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
				fmt.Println("Loop") // Loopが2回出力される
		}

		// 条件付きfor文
		var point int = 0
		// pointが10未満の場合ループを繰り返す
		// 条件がtrueだったらループを繰り返す
		// 条件がfalseになったらループを抜ける
		for point < 10 {
				fmt.Println(point) // 0123456789
				point++
		}

		// 古典的for文
		// for 初期化処理; 条件式; 後処理 {}
		for i := 0; i < 10; i++ {
				fmt.Println(i) // 0123456789
		}
		// iを0で初期化
		// 条件式を満たすとき、ループの処理を実行
		// 後処理を実行
		// 条件式を満たさなくなったら、ループを抜ける

		for i := 0; i < 10; i++ {
				if i == 3 {
						// iが3の場合、処理をスキップする
						// continueはfor文の処理をスキップする
						continue
				}
				if i == 6 {
						// iが6の場合、ループを抜ける
						break
				}
				fmt.Println(i) // 01245
		}

		// 配列でfor文
		// 配列の中身を一つずつ取り出して処理を行う
		var arr [3]int = [3]int{1, 2, 3}
		// len(配列名)で配列の要素数を取得できる
		// [i]で配列のインデックス番号を指定できる
		for i := 0; i < len(arr); i++ {
				fmt.Println(arr[i]) // 123
		}

		// 範囲式for文
		// rangeで配列の要素数を取得できる
		// rangeを使って、配列の要素数分ループを回す
		var arr2 [3]int = [3]int{1, 2, 3}
		for i, v := range arr2 {
				// iにはインデックス番号が入る
				// vには配列の値が入る
				fmt.Println(i, v) // 0 1 1 2 2 3
		}

		// 範囲式for文で要素の値だけ取得する
		// インデックス番号は不要な場合、_で省略できる
		var arr3 [3]int = [3]int{1, 2, 3}
		for _, v := range arr3 {
				fmt.Println(v) // 1 2 3
		}

		// 範囲式for文でインデックス番号だけ取得する
		// 要素の値は不要な場合、変数名を省略できる
		var arr4 [3]int = [3]int{1, 2, 3}
		for i := range arr4 {
				fmt.Println(i) // 0 1 2
		}

		// スライスは要素数が可変
		sl := []string{"Python", "PHP", "Go", "Java"}
		for i, v := range sl {
				fmt.Println(i, v) // 0 Python 1 PHP 2 Go 3 Java
		}

		// スライスで要素の値だけ取得する
		sl2 := []string{"Python", "PHP", "Go", "Java"}
		for _, v := range sl2 {
				fmt.Println(v) // Python PHP Go Java
		}

		// スライスでインデックス番号だけ取得する
		var sl3 []string = []string{"Python", "PHP", "Go", "Java"}
		for i := range sl3 {
				fmt.Println(i) // 0 1 2 3
		}

		// mapで範囲式for文
		// mapはkeyとvalueのペアで管理する
		var m map[string]int = map[string]int{"apple": 100, "banana": 200}
		// 上記だとkeyがstring型、valueがint型
		for k, v := range m {
				fmt.Println(k, v) // apple 100 banana 200
		}

		// mapでvalueだけ取得する
		var m1 map[string]int = map[string]int{"apple": 100, "banana": 200}
		for _, v := range m1 {
				fmt.Println(v) // 100 200
		}

		// mapでkeyだけ取得する
		var m2 map[string]int = map[string]int{"apple": 100, "banana": 200}
		for k := range m2 {
				fmt.Println(k) // apple banana
		}
}
