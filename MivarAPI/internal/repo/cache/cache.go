package cache

import (
	"errors"

	"mivar_robot_api/pkg/cache"
)

type Repo struct {
	cache *cache.Cache
}

func New(cache *cache.Cache) *Repo {
	return &Repo{
		cache: cache,
	}
}

func (r *Repo) AddToCache(data []byte, key string) {
	r.cache.Set(key, data)
}

func (r *Repo) GetFromCache(key string) ([]byte, error) {
	data, ok := r.cache.Get(key)
	if !ok {
		return nil, errors.New("model not found")
	}

	return data.([]byte), nil
}
