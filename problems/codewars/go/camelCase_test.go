// Problem link: https://www.codewars.com/kata/517abf86da9663f1d2000003/train/go

package kata

import "testing"

func TestToCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{""},
			want: "",
		},
		{
			args: args{"The_Stealth_Warrior"},
			want: "TheStealthWarrior",
		},
		{
			args: args{"the-Stealth-Warrior"},
			want: "theStealthWarrior",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToCamelCase(tt.args.s); got != tt.want {
				t.Errorf("ToCamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
