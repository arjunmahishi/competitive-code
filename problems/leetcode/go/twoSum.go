// Problem link: https://leetcode.com/problems/two-sum

package leetcode

func twoSum(nums []int, target int) []int {
	var (
		hashTable = map[int]int{} // number : index
	)
	for i, n := range nums {
		if pi, ok := hashTable[n]; ok {
			return []int{pi, i}
		}
		hashTable[target-n] = i
	}
	return nil
}
