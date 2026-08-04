package cache

import (
	"context"
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

func New(ctx context.Context, ttl time.Duration) *Cache {
	c := &Cache{ttl: ttl, m: make(map[string]entry)}
	go c.startCleanup(ctx)
	return c
}

func (c *Cache) startCleanup(ctx context.Context) {
	interval := c.ttl
	if interval < time.Second {
		interval = time.Second // Prevent tight loop if TTL is very small
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()
			c.mu.Lock()
			for k, v := range c.m {
				if now.After(v.exp) {
					delete(c.m, k)
				}
			}
			c.mu.Unlock()
		}
	}
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
