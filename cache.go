package cache

import (
	"time"
)

type valueWithDeadline struct {
	value    string
	deadline *time.Time
}

type Cache struct {
	storage map[string]valueWithDeadline
}

func NewCache() Cache {
	return Cache{storage: make(map[string]valueWithDeadline)}
}

func (cache *Cache) Get(key string) (string, bool) {
	currentTime := time.Now()
	v, ok := cache.storage[key]
	if !ok {
		return "", ok
	}
	if v.deadline != nil && (v.deadline.After(currentTime) || v.deadline.Equal(currentTime)) {
		delete(cache.storage, key)
		return "", false
	}
	return v.value, ok
}

func (cache *Cache) Put(key, value string) {
	cache.storage[key] = valueWithDeadline{value: value, deadline: nil}
}

func (cache Cache) Keys() []string {
	currentTime := time.Now()
	keys := make([]string, 0, len(cache.storage))
	for k, v := range cache.storage {
		if (v.deadline != nil && v.deadline.Before(currentTime)) || v.deadline == nil {
			keys = append(keys, k)
		}
	}
	return keys
}

func (cache *Cache) PutTill(key, value string, deadline time.Time) {
	cache.storage[key] = valueWithDeadline{value: value, deadline: &deadline}
}
