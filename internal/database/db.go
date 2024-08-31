package database

import (
	"iter"
	"sync"

	"github.com/5aradise/go-message/internal/types"
)

type databse struct {
	data map[string]any
	mu   *sync.RWMutex
}

func New() *databse {
	return &databse{
		data: map[string]any{},
		mu:   new(sync.RWMutex),
	}
}

func (db *databse) Set(k string, v any) {
	db.mu.Lock()
	db.data[k] = v
	db.mu.Unlock()
}

func (db *databse) Get(k string) (any, bool) {
	db.mu.RLock()
	v, ok := db.data[k]
	db.mu.RUnlock()
	return v, ok
}

func (db *databse) Delete(key string) {
	db.mu.Lock()
	delete(db.data, key)
	db.mu.Unlock()
}

func (db *databse) Iter() iter.Seq2[string, any] {
	return func(yield func(string, any) bool) {
		db.mu.RLock()
		for k, v := range db.data {
			db.mu.RUnlock()
			if !yield(k, v) {
				return
			}
			db.mu.RLock()
		}
		db.mu.RUnlock()
	}
}

func (db *databse) SetUser(name string, user *types.User) {
	db.Set(name, user)
}

func (db *databse) GetUserByName(name string) (*types.User, bool) {
	uu, ok := db.Get(name)
	if !ok {
		return nil, false
	}

	u, ok := uu.(*types.User)
	if !ok {
		return nil, false
	}

	return u, true
}

func (db *databse) GetUserByKey(key string) (*types.User, bool) {
	for _, uu := range db.Iter() {
		u, ok := uu.(*types.User)
		if !ok {
			continue
		}
		if u.Key == key {
			return u, true
		}
	}
	return nil, false
}

func (db *databse) IterUsers() iter.Seq2[string, *types.User] {
	return func(yield func(string, *types.User) bool) {
		for name, uu := range db.Iter() {
			u, ok := uu.(*types.User)
			if ok {
				if !yield(name, u) {
					return
				}
			}
		}
	}
}

func (db *databse) DeleteUser(name string) {
	db.Delete(name)
}
