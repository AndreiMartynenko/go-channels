package main

import (
	"fmt"
)

/*
+-------------------------------+
|           hchan                |
+-------------------------------+
| buf        | pointer to buffer |  // Points to the buffer (if buffered)
| elemsize   | size of each elem |  // Size of each element in the channel
| closed     | bool              |  // Indicates if the channel is closed
| qcount     | number of elems   |  // Number of elements currently in buffer
| recvq      | queue of recv g's |  // Queue of waiting receiver goroutines
| sendq      | queue of send g's |  // Queue of waiting sender goroutines
| lock       | mutex             |  // Mutex to protect the channel
+-------------------------------+

*/

// Basic channel communication
//func main() {
//	//Create an unbuffered channel
//	ch := make(chan int)
//
//	// We need a goroutine
//
//	go func() {
//		fmt.Println("Sending data to the channel...")
//		ch <- 10
//		fmt.Println("Data sent.")
//	}()
//
//	// Sleep for a bit to allow the goroutine to start
//	time.Sleep(time.Second)
//
//	// Receive the data from the channel
//	value := <-ch
//	fmt.Println("Received data:", value)
//
//}

// Objective: Implement a program that uses buffered channels to store multiple values.
//func main() {
//	// Create a buffered channel that can hold 3 integers.
//	buffChan := make(chan int, 3)
//
//	send3Ints(buffChan)
//
//	// Receive and print the 3 integers from the channel.
//	for i := 0; i < 3; i++ {
//		time.Sleep(time.Second)
//		value := <-buffChan
//		fmt.Println(value)
//	}
//
//}
//
//func send3Ints(ch chan int) {
//	ch <- 1
//	ch <- 2
//	ch <- 3
//}

// Objective: Use channels to synchronize two goroutines.
//func main() {
//
//	done := make(chan struct{})
//
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	go func() {
//		defer wg.Done()
//		fmt.Println("Goroutine 1: is done")
//	}()
//
//	go func() {
//		defer wg.Done()
//		fmt.Println("Goroutine 2: is done")
//	}()
//
//	wg.Wait()
//	close(done)
//	fmt.Println("All goroutines are done")
//
//}

//Objective: Implement the fan-in pattern where multiple goroutines send data into a single channel.

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Hello"
	}()

	go func() {
		ch2 <- "World"
	}()

	merged := fanIn(ch1, ch2)

	for i := 0; i < 2; i++ {
		fmt.Println(<-merged)
	}
}

func fanIn(input1, input2 chan string) chan string {
	output := make(chan string)
	go func() {
		for {
			output <- <-input1
		}
	}()
	go func() {
		for {
			output <- <-input2
		}
	}()
	return output
}
