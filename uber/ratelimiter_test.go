package uber

import (
	"container/list"
	"reflect"
	"testing"
	"time"
)

func TestNewSlidingWindowRateLimiter(t *testing.T) {
	type args struct {
		requestPerMinute int
	}
	tests := []struct {
		name    string
		args    args
		want    *SlidingWindowRateLimiter
		wantErr bool
	}{
		{"Case when requestPerMinute isn't passed", args{requestPerMinute: -1}, nil, true},
		{"Case when requestPerMinute is passed correctly", args{requestPerMinute: 500}, &SlidingWindowRateLimiter{500, make(map[string]*list.List)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSlidingWindowRateLimiter(tt.args.requestPerMinute)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSlidingWindowRateLimiter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSlidingWindowRateLimiter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlidingWindowRateLimiter_Accept(t *testing.T) {

	rl, err := NewSlidingWindowRateLimiter(3)
	t1 := time.Date(2022, 10, 14, 3, 0, 0, 100, time.UTC)
	t2 := t1.Add(1*time.Minute + 5*time.Second)
	t3 := t1.Add(1*time.Minute + 20*time.Second)
	t4 := t1.Add(1*time.Minute + 45*time.Second)
	t5 := t1.Add(1*time.Minute + 50*time.Second)
	t6 := t1.Add(2*time.Minute + 10*time.Second)
	if err != nil {
		t.Errorf("couldn't create rate limiter instance")
	}
	type args struct {
		clientId string
		timeAt   time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{"XYZ", t1}, true},
		{"t2", args{"XYZ", t2}, true},
		{"t3", args{"XYZ", t3}, true},
		{"t4", args{"XYZ", t4}, true},
		{"t5", args{"XYZ", t5}, false},
		{"t6", args{"XYZ", t6}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rl.Accept(tt.args.clientId, tt.args.timeAt); got != tt.want {
				t.Errorf("SlidingWindowRateLimiter.Accept() = %v, want %v", got, tt.want)
			}
		})
	}
}
