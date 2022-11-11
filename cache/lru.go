package cache

import (
	"container/list"
	"errors"
)

type EvictCallback func(key interface{}, val interface{})

type ICache interface {
	Purge()
	Set(key, val interface{}) bool
	Get(key interface{}) (interface{}, bool)
	Delete(key interface{}) bool
	Contains(key interface{}) bool
	Peek(key interface{}) (interface{}, bool)
	Len() int
	RemoveOldest() (interface{}, interface{}, bool)
	GetOldest() (interface{}, interface{}, bool)
	Keys() []interface{}
}

type LRUCache struct {
	size int
	evictList *list.List
	items map[interface{}]*list.Element
	onEvict EvictCallback
}

type entry struct {
	key interface{}
	value interface{}
}

func NewLRUCache(size int, callback EvictCallback) (*LRUCache, error) {
	if size <= 0 {
		return nil, errors.New("must provide a positive number")
	}

	c := &LRUCache{
		size: size,
		evictList: list.New(),
		items : make(map[interface{}]*list.Element),
		onEvict: callback,
	}

	return c, nil
}


func (c *LRUCache) Purge() {
	for k,v := range c.items {
		if c.onEvict != nil {
			c.onEvict(k, v.Value.(*entry).value)
		}
		delete(c.items, k)
	}

	c.evictList.Init()
}

// Returns true if and only if there was an eviction
func (c *LRUCache) Set(key, val interface{}) bool {
	// check for existing items
	if ent, ok := c.items[key];ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).value = val
		return false
	}

	//Add new items
	ent := &entry{key,val}
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry

	evict := c.evictList.Len() > c.size

	if evict {
		c.removeOldest()
	}

	return evict
}

func (c *LRUCache) Get(key interface{}) (interface{}, bool) {
	if ent,ok := c.items[key];ok {
		c.evictList.MoveToFront(ent)
		if ent.Value.(*entry) == nil {
			return nil, false
		}
		return ent.Value.(*entry).value, false
	}
	return nil, false
}

func (c *LRUCache) Delete(key interface{}) bool {
	if ent, ok := c.items[key];ok {
		c.removeElement(ent)
		return true
	}
	return false
}

func (c *LRUCache) Contains(key interface{}) bool {
	_,ok := c.items[key]
	return ok
}

func (c *LRUCache) Peek(key interface{}) (interface{}, bool) {
	if ent, ok := c.items[key];ok {
		return ent.Value.(*entry).value, true
	}
	return nil, false
}

func (c *LRUCache) Len() int {
	return c.evictList.Len()
}

func (c *LRUCache) RemoveOldest() (interface{}, interface{}, bool) {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}

func (c *LRUCache) GetOldest() (interface{}, interface{}, bool) {
	ent := c.evictList.Back()
	if ent != nil {
		kv := ent.Value.(*entry)
		return kv.key, kv.value, true
	}
	return nil, nil, false
}

func (c *LRUCache) Keys() []interface{} {
	keys := make([]interface{}, 0)
	for ent := c.evictList.Back();ent != nil;ent = ent.Prev() {
		keys = append(keys, ent.Value.(*entry).key)
	}

	return keys
}

func (c *LRUCache) removeOldest() {
	ent := c.evictList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

func (c *LRUCache) removeElement(e *list.Element) {
	c.evictList.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)
	if c.onEvict != nil {
		c.onEvict(kv.key, kv.value)
	}
}




