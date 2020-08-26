// Problem link: https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go

package kata

import "strings"

func ToCamelCase(s string) string {
	if len(s) < 1 {
		return ""
	}
	return strings.Replace(
		s[0:1]+strings.Title(strings.Replace(strings.Replace(s, "_", " ", -1), "-", " ", -1))[1:],
		" ", "", -1,
	)
}
