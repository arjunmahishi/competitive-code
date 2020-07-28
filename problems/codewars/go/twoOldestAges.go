// Link: https://www.codewars.com/kata/511f11d355fe575d2c000001/train/go

package kata

import "sort"

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)
	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}
