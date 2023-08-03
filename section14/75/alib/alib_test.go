package alib

import (
		"testing"
)

var Debug bool = false

func TestAverage(t *testing.T) {
		if Debug {
				t.Skip("スキップします")
		}

		v := Average([]int{1, 2, 3, 4, 5})
		if v != 3 {
				t.Errorf("%v != %v", v, 3)
		}
}

// mainパッケージのディレクトリで go test ./alib でテストを実行する
// 相対パスで指定する
// go test -v ./alib でテストの詳細が表示される
