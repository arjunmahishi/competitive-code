// Link: https://www.codewars.com/kata/55466989aeecab5aac00003e/train/go

package kata

import (
	"reflect"
	"testing"
)

func TestSquaresInRect(t *testing.T) {
	type args struct {
		lng  int
		wdth int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{lng: 5, wdth: 3},
			want: []int{3, 2, 1, 1},
		},
		{
			args: args{lng: 20, wdth: 14},
			want: []int{14, 6, 6, 2, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SquaresInRect(tt.args.lng, tt.args.wdth); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SquaresInRect() = %v, want %v", got, tt.want)
			}
		})
	}
}
