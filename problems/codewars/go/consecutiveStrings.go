// Problem link: https://www.codewars.com/kata/56a5d994ac971f1ac500003e/train/go

package kata

func LongestConsec(strarr []string, k int) string {
	l, s := -1, ""
	for i := 0; i < len(strarr)-(k-1); i++ {
		tmp := consec(strarr, k, i)
		if len(tmp) > l {
			s, l = tmp, len(tmp)
		}
	}
	return s
}

func consec(strarr []string, k, i int) string {
	s := ""
	for j := i; j < i+k; j++ {
		s += strarr[j]
	}
	return s
}
