package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mt sync.Mutex
	mp map[string]string
}

func New() *SafeMap {
	return &SafeMap{
		mp: make(map[string]string),
	}
}

func (m *SafeMap) Set(k, v string) {
	m.mt.Lock()
	defer m.mt.Unlock()
	m.mp[k] = v
}

func (m *SafeMap) Get(k string) (v string, ok bool) {
	m.mt.Lock()
	defer m.mt.Unlock()
	v, ok = m.mp[k]
	return v, ok
}

func TestSafeMap(num int, sm *SafeMap, wg *sync.WaitGroup) {
	key := fmt.Sprintf("key_%d", num)
	value := fmt.Sprintf("value_%d", num)
	sm.Set(key, value)
	wg.Done()
}

func main() {
	sf := New()
	wg := sync.WaitGroup{}

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go TestSafeMap(i+1, sf, &wg)
	}

	wg.Wait()
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key_%d", i+1)
		value, ok := sf.Get(key)

		if ok {
			fmt.Printf("Key(%d) - Value(%s)\n", i+1, value)
		}
	}
}
