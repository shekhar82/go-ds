package systemdesign

import "testing"

func TestLargestNumber(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test case 1", args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}, "9876543210"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestNumber(tt.args.nums); got != tt.want {
				t.Errorf("LargestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
