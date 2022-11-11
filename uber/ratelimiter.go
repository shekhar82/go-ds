package uber

import (
	"container/list"
	"errors"
	"time"
)

type Ratelimiter interface {
	Accept(clientId string, timeAt time.Time) bool
}

type SlidingWindowRateLimiter struct {
	ReqPerMinute int
	Requests     map[string]*list.List
}

func NewSlidingWindowRateLimiter(requestPerMinute int) (*SlidingWindowRateLimiter, error) {
	if requestPerMinute <= 0 {
		return nil, errors.New("request-per-minute can't be <= 0")
	}
	return &SlidingWindowRateLimiter{
		ReqPerMinute: requestPerMinute,
		Requests:     make(map[string]*list.List),
	}, nil
}
func (rl *SlidingWindowRateLimiter) Accept(clientId string, timeAt time.Time) bool {
	now := timeAt.Unix()
	if list, ok := rl.Requests[clientId]; ok && list != nil {
		if list.Back() != nil && now-list.Back().Value.(int64) >= 60 {
			rl.Requests[clientId].Init()
			rl.Requests[clientId].PushFront(now)
			return true
		} else {
			//Remove all access which has timestamp less than 1 minute from current access
			front := list.Front()
			for front != nil && front.Value.(int64) <= (now-60) {
				front = front.Next()
				list.Remove(list.Front())
			}
			if list.Len() < rl.ReqPerMinute {
				list.PushBack(now)
				if now < list.Back().Value.(int64) {
					list.MoveBefore(list.Back(), list.Back().Prev())
				}
				return true
			}
		}
		return false
	}

	rl.Requests[clientId] = list.New()
	rl.Requests[clientId].PushFront(now)
	return true
}

type SlidingWindowCounterRateLimiter struct {
	ReqPerMinute int
	Requests     map[string]*list.List
}

type RequestInstance struct {
	TimeAt  int64
	Counter int
}

func NewSlidingWindowCounterRateLimiter(requestPerMinute int) (*SlidingWindowCounterRateLimiter, error) {
	if requestPerMinute <= 0 {
		return nil, errors.New("request-per-minute can't be <= 0")
	}
	return &SlidingWindowCounterRateLimiter{
		ReqPerMinute: requestPerMinute,
		Requests:     make(map[string]*list.List),
	}, nil
}

func (rl *SlidingWindowCounterRateLimiter) len(clientId string) int {
	if list, ok := rl.Requests[clientId]; ok {
		totalLength := 0
		for front := list.Front(); front != nil; front = front.Next() {
			totalLength += front.Value.(RequestInstance).Counter
		}
		return totalLength
	}
	return -1
}

func (rl *SlidingWindowCounterRateLimiter) Accept(clientId string, timeAt time.Time) bool {
	now := timeAt.Unix()
	if list, ok := rl.Requests[clientId]; ok {
		if now-list.Back().Value.(RequestInstance).TimeAt >= 60 {
			requestInstance := RequestInstance{now, 1}
			list.Init()
			list.PushFront(requestInstance)
		} else {
			front := list.Front()
			for front != nil && front.Value.(RequestInstance).TimeAt <= (now-60) {
				front = front.Next()
				list.Remove(list.Front())
			}

			if rl.len(clientId) != -1 && rl.len(clientId) < rl.ReqPerMinute {
				if now > list.Back().Value.(RequestInstance).TimeAt {
					requestInstance := RequestInstance{now, 1}
					list.PushBack(requestInstance)
				} else {
					for back := list.Back(); back != nil && back.Value.(RequestInstance).TimeAt >= now; back = back.Prev() {
						if back.Value.(RequestInstance).TimeAt == now {
							rc := back.Value.(RequestInstance)
							rc.Counter += 1
							back.Value = rc
						} else {
							rc := RequestInstance{now, 1}
							list.InsertAfter(rc, back)
						}
					}
				}
				return true
			}
			return false
		}
	}

	requestInstance := RequestInstance{now, 1}

	rl.Requests[clientId] = list.New()
	rl.Requests[clientId].PushFront(requestInstance)
	return true
}
