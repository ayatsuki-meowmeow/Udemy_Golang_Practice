package alib

// Average は、スライスの平均値を返す
func Average(s []int) int {
		// totalには、スライスの合計値が入る
		total := 0
		for _, i := range s {
				total += i
		}
		// totalをスライスの要素数で割る
		return int(total / len(s))
}
