package main

import (
		"fmt"
		"time"
)

// time

func main() {
		// 時間の間隔を表すDuration型
		// time.Duration型はint64型の別名型
		fmt.Println(time.Hour) // 1h0m0s
		fmt.Printf("%T\n", time.Hour) // time.Duration
		fmt.Println(time.Minute) // 1m0s
		fmt.Println(time.Second) // 1s
		fmt.Println(time.Millisecond) // 1ms
		fmt.Println(time.Microsecond) // 1µs
		fmt.Println(time.Nanosecond) // 1ns

		// time.Duration型を文字列から生成する
		// time.ParseDuration()で文字列からtime.Duration型を生成する
		// time.ParseDuration(文字列)
		// 文字列は1h30mのように時間、分、秒を表す単位をつけて表す
		d, _ := time.ParseDuration("2h30m")
		fmt.Println(d) // 2h30m0s

		// time.Duration型とtime.Time型の演算
		// 現在時刻の取得
		t3 := time.Now()
		// 現在時刻から2分15秒後の時刻を取得する
		t3 = t3.Add(2*time.Minute + 15*time.Second)
		fmt.Println(t3)
}
