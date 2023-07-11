package main

// スライス
// 可変長引数

import "fmt"

// Sum関数には引数をいくつでも渡せる
func Sum(s ...int) int { // 可変長引数は型名の前に...をつける
		n := 0
		for _, v := range s {
			n += v
		}
		return n
}

// SumとSum2は同じ
func Sum2(s ...int) (n int) {
	for _, v := range s {
		n += v
	}
	return
}

// int型以外の可変長引数もできる
func Sum3(s ...string) (n string) {
	for _, v := range s {
		n += v
	}
	return
}

func main() {
		fmt.Println(Sum(1, 2)) // 3
		fmt.Println(Sum(1, 2, 3, 4, 5)) // 15
		fmt.Println(Sum()) // 0

		// スライスを渡すこともできる
		var sl []int = []int{1, 2, 3, 4, 5, 6, 7, 8 ,9, 10}
		fmt.Println(Sum(sl...)) // 55
		fmt.Println(Sum2(sl...)) // 55

		var sl2 []string = []string{"月", "が", "綺麗", "です", "ね"}
		fmt.Println(Sum3(sl2...)) // 月が綺麗ですね
}
