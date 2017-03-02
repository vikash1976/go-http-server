package store

import "sync"

type Store struct {
mu sync.RWMutex
m  map[string]string
}

func (s *Store) Set(key, val string) {
s.mu.Lock()
defer s.mu.Unlock()
if s.m == nil {
s.m = make(map[string]string)
}
s.m[key] = val
}

func (s *Store) Get(key string) (string, bool) {
s.mu.RLock()
defer s.mu.RUnlock()
v, ok := s.m[key]
return v, ok
}

var defaultStore Store

func Set(key, val string) {
defaultStore.Set(key, val)
}

func Get(key string) (string, bool) {
return defaultStore.Get(key)
}
