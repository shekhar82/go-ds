package ds

import (
	"reflect"
	"testing"
)

func TestGetAQueue(t *testing.T) {
	tests := []struct {
		name string
		want *Queue
	}{
		{"Test to find if this method returns properly initialised empty queue", &Queue{Elements: make([]interface{}, 0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAQueue(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Enqueue(t *testing.T) {
	type fields struct {
		Elements []interface{}
	}
	type args struct {
		element interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"test if uninitialised queue should throw errors", fields{Elements: nil}, args{element: 1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Elements: tt.fields.Elements,
			}
			if err := q.Enqueue(tt.args.element); (err != nil) != tt.wantErr {
				t.Errorf("Queue.Enqueue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestQueue_IsEmpty(t *testing.T) {
	type fields struct {
		Elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Elements: tt.fields.Elements,
			}
			if got := q.IsEmpty(); got != tt.want {
				t.Errorf("Queue.IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Size(t *testing.T) {
	type fields struct {
		Elements []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Elements: tt.fields.Elements,
			}
			if got := q.Size(); got != tt.want {
				t.Errorf("Queue.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Peek(t *testing.T) {
	type fields struct {
		Elements []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Elements: tt.fields.Elements,
			}
			got, err := q.Peek()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Peek() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Peek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueue_Dequeue(t *testing.T) {
	type fields struct {
		Elements []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &Queue{
				Elements: tt.fields.Elements,
			}
			got, err := q.Dequeue()
			if (err != nil) != tt.wantErr {
				t.Errorf("Queue.Dequeue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Queue.Dequeue() = %v, want %v", got, tt.want)
			}
		})
	}
}
