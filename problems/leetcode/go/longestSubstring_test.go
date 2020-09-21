// Problem link: https://leetcode.com/problems/longest-substring-without-repeating-characters/

package leetcode

import (
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1 in LeetCode",
			args: args{"abcabcbb"},
			want: 3,
		},
		{
			name: "Example 2 in LeetCode",
			args: args{"bbbbb"},
			want: 1,
		},
		{
			name: "Example 3 in LeetCode",
			args: args{"pwwkew"},
			want: 3,
		},
		{
			name: "Example 4 in LeetCode",
			args: args{""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
