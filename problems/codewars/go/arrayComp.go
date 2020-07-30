// Link: https://www.codewars.com/kata/550498447451fbbd7600041c/train/go

package kata

import "sort"

func Comp(array1 []int, array2 []int) bool {
	if array1 == nil || array2 == nil || len(array1) != len(array2) {
		return false
	}
	a, b, l := abs(array1), abs(array2), len(array1)
	sort.Ints(a)
	sort.Ints(b)
	if l != 0 && a[l-1] > b[l-1] {
		b, a = a, b
	}

	for i := 0; i < l; i++ {
		if a[i]*a[i] != b[i] {
			return false
		}
	}
	return true
}

func abs(arr []int) []int {
	newArr := []int{}
	for _, e := range arr {
		if e < 0 {
			e *= -1
		}
		newArr = append(newArr, e)
	}
	return newArr
}
