package systemdesign

import "testing"

func TestMinimumTime(t *testing.T) {
	type args struct {
		time       []int
		totalTrips int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"First test from leet code", args{time: []int{1, 2, 3}, totalTrips: 5}, 3},
		{"second test from leet code", args{time: []int{2}, totalTrips: 1}, 2},
		{"failed test from leet code", args{time: []int{3, 3, 8}, totalTrips: 6}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumTime(tt.args.time, tt.args.totalTrips); got != tt.want {
				t.Errorf("MinimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
