// Problem link: https://www.codewars.com/kata/550498447451fbbd7600041c/train/go

package kata

import "testing"

func TestComp(t *testing.T) {
	type args struct {
		array1 []int
		array2 []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{
				array1: []int{121, 144, 19, 161, 19, 144, 19, 11},
				array2: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			want: true,
		},
		{
			args: args{
				array1: []int{121, 144, 19, -161, 19, -144, 19, -11},
				array2: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			want: true,
		},
		{
			args: args{
				array1: []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
				array2: []int{121, 144, 19, 161, 19, 144, 19, 11},
			},
			want: true,
		},
		{
			args: args{
				array1: []int{121, 144, 19, 161, 19, 144, 19, 11},
				array2: []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			want: false,
		},
		{
			args: args{
				array1: nil,
				array2: []int{11 * 21, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19},
			},
			want: false,
		},
		{
			args: args{
				array1: []int{},
				array2: []int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Comp(tt.args.array1, tt.args.array2); got != tt.want {
				t.Errorf("Comp() = %v, want %v", got, tt.want)
			}
		})
	}
}
