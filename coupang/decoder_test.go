package coupang

import "testing"

func TestDecode(t *testing.T) {
	type args struct {
		encoded string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Empty string", args{encoded: ""}, ""},
		{"String with no open bracket", args{encoded: "abc"}, "abc"},
		{"Proper string 1", args{encoded: "3[a]2[bc]"}, "aaabcbc"},
		{"Proper string 2", args{encoded: "3[a2[c]]"}, "accaccacc"},
		{"Proper string 3", args{encoded: "2[abc]3[cd]ef"}, "abcabccdcdcdef"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.encoded); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
