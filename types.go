package main

import (
	"sync"

	"github.com/dgryski/go-farm"
)

type ConcurrentHashMay struct {
	maps  []map[string]any
	seg   int
	locks []sync.RWMutex
	seed  uint32
}

func NewConcurrentHashMay(seg, cap int) *ConcurrentHashMay {
	maps := make([]map[string]any, seg)
	locks := make([]sync.RWMutex, seg)
	for i := range maps {
		maps[i] = make(map[string]any, cap/seg)
	}

	return &ConcurrentHashMay{maps, seg, locks, 0}
}

func (m *ConcurrentHashMay) getSegIndex(key string) int {
	hash := int(farm.Hash32WithSeed([]byte(key), m.seed))
	return hash % m.seg
}

func (m *ConcurrentHashMay) Set(key string, val any) {
	index := m.getSegIndex(key)
	m.locks[index].Lock()
	defer m.locks[index].Unlock()

	m.maps[index][key] = val
}

func (m *ConcurrentHashMay) Get(key string) (val any, exists bool) {
	index := m.getSegIndex(key)
	m.locks[index].RLock()
	defer m.locks[index].RUnlock()
	val, exists = m.maps[index][key]
	return
}

const (
	a = 1 << iota
	b
	c
	d
)
