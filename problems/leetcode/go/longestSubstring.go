// Problem link: https://leetcode.com/problems/longest-substring-without-repeating-characters/

package leetcode

func lengthOfLongestSubstring(s string) int {
	count := 0
	for i := 0; i <= len(s); i++ {
		if maxCount := maxUnique(s[i:]); maxCount > count {
			count = maxCount
		}
	}
	return count
}

func maxUnique(s string) int {
	hm, count := map[rune]bool{}, 0
	for _, c := range s {
		if _, ok := hm[c]; ok {
			return count
		}
		hm[c] = true
		count++
	}
	return count
}
