package uber

import (
	"container/list"
	"hash/fnv"
	"math/rand"
)

type entry struct {
	key   string
	value interface{}
}

type kvCell struct {
	size          int
	collisionList *list.List
}

type KVStore struct {
	Capacity int
	Cells    []*kvCell
}

func NewKVStore(capacity int) KVStore {
	return KVStore{
		Capacity: capacity,
		Cells:    make([]*kvCell, capacity),
	}
}

func (kv *KVStore) Set(key string, val interface{}) {
	indexForKey := calculateIdx(key, kv.Capacity)

	if kv.Cells[indexForKey] != nil {
		front := kv.Cells[indexForKey].collisionList.Front()
		for front != nil {
			if front.Value.(*entry).key == key {
				front.Value.(*entry).value = val
				break
			}
			front = front.Next()
		}

		if front == nil {
			entry := entry{key, val}
			kv.Cells[indexForKey].collisionList.PushBack(entry)
			kv.Cells[indexForKey].size += 1
		}
		return
	}

	entry := entry{key, val}
	kv.Cells[indexForKey] = new(kvCell)
	kv.Cells[indexForKey].size = 1
	kv.Cells[indexForKey].collisionList = list.New()
	kv.Cells[indexForKey].collisionList.PushFront(entry)

}

func (kv *KVStore) RandomKey() string {
	randomIdx := rand.Intn(kv.Capacity) - 1
}

func calculateIdx(key string, capacity int) int {
	if len(key) == 0 {
		return -1
	}
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32())%capacity - 1
}
