// Problem link: https://leetcode.com/problems/string-to-integer-atoi/

package leetcode

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{
			name: "Example 1 from LeetCode",
			str:  "42",
			want: 42,
		},
		{
			name: "example 1 from leetcode",
			str:  "    -42",
			want: -42,
		},
		{
			name: "example 1 from leetcode",
			str:  "4193 with words",
			want: 4193,
		},
		{
			name: "example 1 from leetcode",
			str:  "words and 987",
			want: 0,
		},
		{
			name: "example 1 from leetcode",
			str:  "-91283472332",
			want: -2147483648,
		},
		{
			name: "example 1 from leetcode",
			str:  "+1",
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.str); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
