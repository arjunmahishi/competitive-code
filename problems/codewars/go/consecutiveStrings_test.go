// Link: https://www.codewars.com/kata/56a5d994ac971f1ac500003e/train/go

package kata

import "testing"

func TestLongestConsec(t *testing.T) {
	type args struct {
		strarr []string
		k      int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				strarr: []string{"zone", "abigail", "theta", "form", "libe", "zas"},
				k:      2,
			},
			want: "abigailtheta",
		},
		{
			args: args{
				strarr: []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
				k:      1,
			},
			want: "oocccffuucccjjjkkkjyyyeehh",
		},
		{
			args: args{
				strarr: []string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"},
				k:      2,
			},
			want: "wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestConsec(tt.args.strarr, tt.args.k); got != tt.want {
				t.Errorf("LongestConsec() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}
