package main

import (
	"fmt"
	"time"
)

/**
 * Thread example
 * Problem: In this example we are competing for memory space
//  */
// func counter(n int) {
// 	for i := range n {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	go counter(5)
// 	go counter(5)
// 	counter(5)
// }

/**
 * Channel example
 * Shares data between threads without competing
 * for memory space
 */
func worker(workerID int, data chan int) {
	for i := range data {
		fmt.Printf("worker %d received %d\n", workerID, i)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)

	// 3 workers will share the same channel
	go worker(1, ch)
	go worker(2, ch)
	go worker(3, ch)

	for i := range 10 {
		ch <- i
	}
}
