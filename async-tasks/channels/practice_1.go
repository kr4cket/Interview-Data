package main

import (
	"fmt"
	"sync"
	"time"
)

func inputChannel(num int, channel chan int, wg *sync.WaitGroup) {
	select {
	case <-channel:
		channel <- num
		fmt.Println("Func completed")
	default:
		fmt.Println("Channel close")
	}
}

func testCommonData(channel chan int, goNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	temp := <-channel
	channel <- goNum
	for i := 0; i < 5; i++ {
		fmt.Printf("It's goroutine: %d\n", goNum)
		fmt.Printf("Got data from goroutine: %d\n", temp)
		fmt.Printf("Data changed!\n")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	wg := sync.WaitGroup{}
	channel := make(chan int, 1)
	channel <- 1
	defer close(channel)
	wg.Add(2)

	go testCommonData(channel, 1, &wg)
	go testCommonData(channel, 2, &wg)

	wg.Wait()
	fmt.Println(fmt.Sprintf("Func finished"))
}
