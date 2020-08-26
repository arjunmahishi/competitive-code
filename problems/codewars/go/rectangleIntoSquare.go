// Problem link: https://www.codewars.com/kata/55466989aeecab5aac00003e/train/go

package kata

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func SquaresInRect(lng int, wdth int) []int {
	if lng == wdth {
		return nil
	}
	var (
		sq   = []int{}
		n, m = minMax(lng, wdth)
	)
	for n > 0 && m > 0 {
		sq = append(sq, n)
		n, m = minMax(n, m-n)
	}
	return sq
}
