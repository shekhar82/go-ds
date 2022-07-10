package ds

import "errors"

type Queue struct {
	Elements []interface{}
}

func GetAQueue() *Queue {
	return &Queue{
		Elements: make([]interface{}, 0),
	}
}

func (q *Queue) Enqueue(element interface{}) error {
	if q.Elements != nil {
		q.Elements = append(q.Elements, element)
		return nil
	}

	return errors.New("uninitialised queue. please use GetAQueue() method to get queue instance")
}

func (q *Queue) IsEmpty() bool {
	return q.Elements != nil && len(q.Elements) == 0
}

func (q *Queue) Size() int {
	if q.Elements != nil {
		return len(q.Elements)
	}
	return 0
}

func (q *Queue) Peek() (interface{}, error) {
	if q.Elements != nil {
		return q.Elements[0], nil
	}

	return nil, errors.New("uninitialised queue. please use GetAQueue() method to get queue instance")
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.Elements != nil {
		element := q.Elements[0]
		q.Elements = q.Elements[1:]
		return element, nil
	}

	return nil, errors.New("uninitialised queue. please use GetAQueue() method to get queue instance")
}
