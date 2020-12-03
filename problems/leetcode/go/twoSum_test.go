// Problem link: https://leetcode.com/problems/two-sum

package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1 on leet code",
			args: args{nums: []int{2, 7, 11, 15}, target: 9},
			want: []int{0, 1},
		},
		{
			name: "Example 2 on leet code",
			args: args{nums: []int{3, 2, 4}, target: 6},
			want: []int{1, 2},
		},
		{
			name: "Example 3 on leet code",
			args: args{nums: []int{3, 3}, target: 6},
			want: []int{0, 1},
		},
		{
			name: "Example 4 on leet code",
			args: args{nums: []int{4, 3}, target: 6},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
