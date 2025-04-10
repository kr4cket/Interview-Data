package main

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mt      sync.Mutex
	counter int
}

func (s *SafeCounter) Counter() int {
	s.mt.Lock()
	defer s.mt.Unlock()
	return s.counter
}

func (s *SafeCounter) Inc() {
	s.mt.Lock()
	defer s.mt.Unlock()
	s.counter++
}

func safeValueTest(sc *SafeCounter, wg *sync.WaitGroup) {
	for i := 0; i < 1000; i++ {
		sc.Inc()
	}

	wg.Done()
}

// Безопасный доступ к общей переменной
func main() {
	cntr := &SafeCounter{
		counter: 0,
	}

	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go safeValueTest(cntr, &wg)
	}

	wg.Wait()
	fmt.Println(cntr.Counter()) // Результат должен быть равен 10000
}
