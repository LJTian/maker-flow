package cache

import (
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

type entry struct {
	value any
	exp   time.Time
}

type Cache struct {
	ttl time.Duration
	mu  sync.RWMutex
	m   map[string]entry
	sf  singleflight.Group
}

func New(ttl time.Duration) *Cache {
	return &Cache{ttl: ttl, m: make(map[string]entry)}
}

func (c *Cache) GetOrLoad(key string, load func() (any, error)) (any, error, bool) {
	c.mu.RLock()
	if e, ok := c.m[key]; ok && time.Now().Before(e.exp) {
		c.mu.RUnlock()
		return e.value, nil, true
	}
	c.mu.RUnlock()

	v, err, _ := c.sf.Do(key, func() (any, error) {
		c.mu.RLock()
		if e, ok := c.m[key]; ok && time.Now().Before(e.exp) {
			c.mu.RUnlock()
			return e.value, nil
		}
		c.mu.RUnlock()

		val, err := load()
		if err != nil {
			return nil, err
		}
		c.mu.Lock()
		c.m[key] = entry{value: val, exp: time.Now().Add(c.ttl)}
		c.mu.Unlock()
		return val, nil
	})
	return v, err, false
}
