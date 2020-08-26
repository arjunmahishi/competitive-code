// Problem link: https://www.codewars.com/kata/511f11d355fe575d2c000001/train/go

package kata

import (
	"reflect"
	"testing"
)

func TestTwoOldestAges(t *testing.T) {
	type args struct {
		ages []int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			args: args{
				ages: []int{6, 5, 83, 5, 3, 18},
			},
			want: [2]int{18, 83},
		},
		{
			args: args{
				ages: []int{1, 5, 87, 45, 8, 8},
			},
			want: [2]int{45, 87},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TwoOldestAges(tt.args.ages); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TwoOldestAges() = %v, want %v", got, tt.want)
			}
		})
	}
}
