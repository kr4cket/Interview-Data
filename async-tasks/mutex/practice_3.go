package main

import (
	"fmt"
	"sync"
)

type SafeMapRW struct {
	rwmt sync.RWMutex
	mp   map[int]string
}

func NewSafeMapRW() *SafeMapRW {
	return &SafeMapRW{
		rwmt: sync.RWMutex{},
		mp:   make(map[int]string),
	}
}

func (sf *SafeMapRW) Read(key int) (string, bool) {
	sf.rwmt.RLock()
	defer sf.rwmt.RUnlock()
	v, ok := sf.mp[key]
	return v, ok
}

func (sf *SafeMapRW) Write(key int, value string) {
	sf.rwmt.Lock()
	defer sf.rwmt.Unlock()
	sf.mp[key] = value
}

func setData(num int, wg *sync.WaitGroup, sf *SafeMapRW) {
	defer wg.Done()
	sf.Write(num, fmt.Sprintf("Val -> %d", num))
}

func getData(num int, wg *sync.WaitGroup, sf *SafeMapRW) {
	defer wg.Done()
	v, ok := sf.Read(num)
	if ok {
		fmt.Printf("Key -> %d, %s \n", num, v)
	}
}

func main() {
	sf := NewSafeMapRW()
	wg := sync.WaitGroup{}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go setData(i, &wg, sf)
	}
	wg.Wait()

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go getData(i, &wg, sf)
	}
	wg.Wait()
}
