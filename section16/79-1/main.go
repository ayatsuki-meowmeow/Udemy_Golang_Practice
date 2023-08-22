package main

import (
		"fmt"
		"math"
)

// math
// 数学に関連した機能がまとめられたパッケージ

func main() {
		// 数学的な定数
		// 円周率
		fmt.Println(math.Pi) // 3.141592653589793

		// 2の平方根
		fmt.Println(math.Sqrt2) // 1.4142135623730951
		// nの平方根は、math.Sqrt(n)、もしくは、math.Sqrtn で取得できる

		// 数値型に関する定数
		// float32で表現可能な最大値
		fmt.Println(math.MaxFloat32) // 3.4028234663852886e+38
		// float32で表現可能な0ではない最小値
		fmt.Println(math.SmallestNonzeroFloat32) // 1.401298464324817e-45
		// float64で表現可能な最大値
		fmt.Println(math.MaxFloat64) // 1.7976931348623157e+308
		// float64で表現可能な0ではない最小値
		fmt.Println(math.SmallestNonzeroFloat64) // 5e-324
		// int64で表現可能な最大値
		fmt.Println(math.MaxInt64) // 9223372036854775807
		// int64で表現可能な最小値
		fmt.Println(math.MinInt64) // -9223372036854775808
		// int32で表現可能な最大値
		fmt.Println(math.MaxInt32) // 2147483647
		// int32で表現可能な最小値
		fmt.Println(math.MinInt32) // -2147483648

		// この他にも、各種整数型についても、最大値と最小値を取得できる
		// 例えば、uint8で表現可能な最大値は、math.MaxUint8 で取得できる
		// また、int8で表現可能な最小値は、math.MinInt8 で取得できる
		fmt.Println(math.MaxUint8) // 255
		fmt.Println(math.MinInt8) // -128
}
