package main

import (
		"fmt"
    "time"
)

// time
// 時刻の比較

func main() {
		// 時刻の差分を取得
		t5 := time.Date(2020, 4, 6, 10, 10, 10, 0, time.Local)
		var t6 time.Time = time.Now()
		fmt.Println(t6)

		// t5 - t6
		// time.Sub()で時刻の差分を取得する
		// time.Sub(時刻)
		d2 := t5.Sub(t6)
		fmt.Println(d2) // ~h~m~s

		// 時刻の比較
		// time.Time型の前後関係を判定する
		// 結果はbool型で返される
		// time.Before(time2) で timeがtime2より前の時刻かどうかを判定する
		// time.After(time2) で timeがtime2より後の時刻かどうかを判定する
		fmt.Println(t6.Before(t5)) // false
		fmt.Println(t6.After(t5)) // true
		fmt.Println(t5.Before(t6)) // true
		fmt.Println(t5.After(t6)) // false
}
