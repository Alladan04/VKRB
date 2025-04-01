package cache

import (
	"errors"
)

type Repo struct {
	modelCache    Cache
	labirintCache Cache
}

func New(modelCache Cache, labirintCache Cache) *Repo {
	return &Repo{
		modelCache:    modelCache,
		labirintCache: labirintCache,
	}
}

func (r *Repo) AddToCache(data []byte, key string) {
	r.modelCache.Set(key, data)
}

func (r *Repo) GetFromCache(key string) ([]byte, error) {
	data, ok := r.modelCache.Get(key)
	if !ok {
		return nil, errors.New("model not found")
	}

	return data.([]byte), nil
}

func (r *Repo) AddLabirintToCache(data [][]uint8, key string) {
	r.labirintCache.Set(key, data)
}

func (r *Repo) GetLabirintFromCache(key string) ([][]uint8, error) {
	data, ok := r.labirintCache.Get(key)
	if !ok {
		return nil, errors.New("model not found")
	}

	return data.([][]uint8), nil
}
