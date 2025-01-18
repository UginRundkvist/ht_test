package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache    // Remove me after realization.
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (lc *lruCache) Set(k Key, v interface{}) bool {
	item, ok := lc.items[k]
	if ok {
		c, _ := item.Value.(cacheItem)
		c.val = v
		item.Value = c
		lc.queue.PushFront(c)
		return true
	}
	c := cacheItem{key: k, val: v}
	lc.items[k] = lc.queue.PushFront(c)

	if lc.queue.Len() > lc.capacity {
		b := lc.queue.Back()
		lc.queue.Remove(b)
		d, _ := b.Value.(cacheItem)
		delete(lc.items, d.key)
	}

	return false
}

func (lc *lruCache) Get(k Key) (interface{}, bool) {
	item, ok := lc.items[k]
	if ok {
		c, _ := item.Value.(cacheItem)
		item.Value = cacheItem{key: k, val: item.Value}
		lc.queue.MoveToFront(item)
		return c.val, true
	}
	return nil, false
}

func (lc *lruCache) Clear() {
	lc.queue = nil
	lc.items = map[Key]*ListItem{}
}

type cacheItem struct {
	key Key
	val interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
