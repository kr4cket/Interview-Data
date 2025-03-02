package main

import (
	"fmt"
	"sync"
)

type SafeQueue struct {
	queue []int
	mt    sync.Mutex
}

func NewSafeQueue() *SafeQueue {
	return &SafeQueue{
		queue: make([]int, 0),
	}
}

func (q *SafeQueue) Enqueue(v int) {
	q.mt.Lock()
	defer q.mt.Unlock()
	q.queue = append(q.queue, v)
}

func (q *SafeQueue) Dequeue() int {
	q.mt.Lock()
	defer q.mt.Unlock()
	if len(q.queue) == 0 {
		return -1
	}
	v := q.queue[0]
	q.queue = q.queue[1:]
	return v
}

func setDataToQueue(num int, wg *sync.WaitGroup, queue *SafeQueue) {
	defer wg.Done()
	queue.Enqueue(num)
}

func getDataFromQueue(wg *sync.WaitGroup, queue *SafeQueue) {
	defer wg.Done()
	val := queue.Dequeue()
	fmt.Printf("Val -> %v\n", val)
}

func main() {
	q := NewSafeQueue()
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go setDataToQueue(i, &wg, q)
	}
	wg.Wait()

	wg.Add(11)
	for i := 0; i < 11; i++ {
		go getDataFromQueue(&wg, q)
	}
	wg.Wait()
}
