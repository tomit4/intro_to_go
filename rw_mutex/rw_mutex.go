package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type safeCounter struct {
	counts map[string]int
	mux    *sync.RWMutex
}

func (sc safeCounter) inc(key string) {
	sc.mux.Lock()
	defer sc.mux.Unlock()
	sc.slowIncrement(key)
}

func (sc safeCounter) val(key string) int {
	sc.mux.RLock()
	defer sc.mux.RUnlock()
	return sc.counts[key]
}
