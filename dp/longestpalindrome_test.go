package dp

import "testing"

func TestLongestPalindrome(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single string", args{input: "a"}, "a"},
		{"2 same chars in string", args{input: "bb"}, "bb"},
		{"2 different chars in string", args{input: "ba"}, "b"},
		{"multiple chars test 1", args{input: "babad"}, "aba"},
		{"multiple chars test 1", args{input: "cbbdracecar"}, "racecar"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LongestPalindrome(tt.args.input); got != tt.want {
				t.Errorf("LongestPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
