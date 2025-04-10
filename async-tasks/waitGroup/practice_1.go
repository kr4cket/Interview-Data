package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("It's a %d goroutine, value: %d \n", number, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var gos int
	var wg sync.WaitGroup
	fmt.Print("Enter integer value: ")

	_, err := fmt.Scanf("%d", &gos)
	if err != nil {
		fmt.Println(err)
		return
	}

	wg.Add(gos)

	for i := 0; i < gos; i++ {
		go printNumbers(i+1, &wg)
	}

	wg.Wait()
	fmt.Println("Main function finished")
}
