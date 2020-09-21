// Problem link: https://leetcode.com/problems/string-to-integer-atoi/

package leetcode

// ASCII values: '-' - 45, '0' - 48, '9' - 57

func myAtoi(str string) int {
	num, sign, metSign, metNumber := 0, 1, false, false
	for _, r := range str {
		switch {
		case r == '-' && !metSign:
			sign = -1
			metSign = true
			continue
		case r == '+' && !metSign:
			metSign = true
		case r >= '0' && r <= '9':
			num *= 10
			num += int(r) - 48
			metNumber = true
			continue
		case r == ' ' && !metNumber:
			continue
		default:
			return limitInt(num * sign)
		}
	}
	return limitInt(num * sign)
}

func limitInt(n int) int {
	if n > 2147483648 {
		return 2147483648
	}
	if n < -2147483648 {
		return -2147483648
	}
	return n
}
