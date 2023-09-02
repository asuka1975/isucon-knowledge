package cache

import (
	"time"
	"sync"
)

type Timespan struct {
	Start 	int64
	Seconds int64
}

type Cache struct {
	m sync.Map
}

type Item struct {
	Value interface{}
	Expire Timespan
}

func (c *Cache) Get(key string, getter func() interface{}, seconds int64) interface{} {
	rawValue, ok := c.m.Load(key)
	if !ok {
		value := getter()
		c.m.Store(key, Item { Value: value, Expire: Timespan{time.Now().UnixNano(), seconds} })
		return value
	}

	item, ok_ := rawValue.(Item)
	value := item.Value
	expire := item.Expire
	if !ok_ || expire.Seconds >= 0 && time.Now().UnixNano() > expire.Start + expire.Seconds * 1000000000 {
		value = getter()
		c.m.Store(key, Item { Value: value, Expire: Timespan{time.Now().UnixNano(), seconds} })
	}

	return value
}

func New() *Cache {
	return &Cache{
		m: sync.Map{},
	}
}