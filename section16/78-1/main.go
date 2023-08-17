package main

import (
		"fmt"
		"time"
)

// time

func main() {
		// 時間の生成
		// time.Now()で現在時刻を取得する
		t := time.Now()
		fmt.Println(t)

		// 指定した時間の生成
		// time.Date()で指定した時間を取得する
		// time.Date(年, 月, 日, 時, 分, 秒, ナノ秒, タイムゾーン)
		t2 := time.Date(2020, 4, 6, 10, 10, 10, 0, time.Local)
		fmt.Println(t2)
		// 年だけ取得する
		fmt.Println(t2.Year())
		// 通算日(1月1日からの日数)だけ取得する
		fmt.Println(t2.YearDay())
		// 月だけ取得する
		fmt.Println(t2.Month())
		// 曜日だけ取得する
		fmt.Println(t2.Weekday())
		// 日だけ取得する
		fmt.Println(t2.Day())
		// 時だけ取得する
		fmt.Println(t2.Hour())
		// 分だけ取得する
		fmt.Println(t2.Minute())
		// 秒だけ取得する
		fmt.Println(t2.Second())
		// ナノ秒だけ取得する
		fmt.Println(t2.Nanosecond())
		// タイムゾーンだけ取得する
		fmt.Println(t2.Zone())
}
