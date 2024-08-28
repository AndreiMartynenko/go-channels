package main

import (
	"fmt"
	"sync"
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

//Objective: Implement the fan-in pattern where multiple
//goroutines send data into a single channel.

//func main() {
//	ch1 := make(chan string)
//	ch2 := make(chan string)
//
//	go func() {
//		ch1 <- "Hello"
//	}()
//
//	go func() {
//		ch2 <- "World"
//	}()
//
//	merged := fanIn(ch1, ch2)
//
//	for i := 0; i < 2; i++ {
//		fmt.Println(<-merged)
//	}
//}
//
//func fanIn(input1, input2 chan string) chan string {
//	output := make(chan string)
//	go func() {
//		for {
//			output <- <-input1
//		}
//	}()
//	go func() {
//		for {
//			output <- <-input2
//		}
//	}()
//	return output
//}

//Objective: Use the select statement to handle multiple channel operations.

//func main() {
//	ch1 := make(chan string)
//	ch2 := make(chan string)
//
//	go func() {
//		time.Sleep(time.Second)
//		ch1 <- "Hello"
//	}()
//
//	go func() {
//		time.Sleep(2 * time.Second)
//		ch2 <- "World"
//	}()
//
//	for i := 0; i < 2; i++ {
//		select {
//		case msg1 := <-ch1:
//			fmt.Println(msg1)
//		case msg2 := <-ch2:
//			fmt.Println(msg2)
//		}
//	}
//}

//Objective: Implement a timeout mechanism using channels and the time.After function.

//func main() {
//	ch := make(chan string)
//
//	go func() {
//		time.Sleep(2 * time.Second)
//		ch <- "Hello"
//	}()
//
//	select {
//	case msg := <-ch:
//		fmt.Println(msg)
//	case <-time.After(1 * time.Second):
//		fmt.Println("Timeout")
//	}
//}

//Objective: Build a simple worker pool using channels to distribute tasks among multiple workers.

//func main() {
//	tasks := make(chan int, 10)
//	results := make(chan int, 10)
//
//	var wg sync.WaitGroup
//	for i := 0; i < 3; i++ { // Create 3 workers
//		wg.Add(1)
//		go worker(tasks, results, &wg)
//	}
//
//	for i := 0; i < 10; i++ {
//		tasks <- i
//	}
//	close(tasks)
//
//	wg.Wait()
//	close(results)
//
//	for result := range results {
//		fmt.Println("Result: ", result)
//	}
//
//}
//
//func worker(tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
//	defer wg.Done()
//	for task := range tasks {
//		results <- task * 2 // Simple processing: multiply task by 2
//	}
//}

// Objective: Understand and avoid deadlocks in Go channels.

//func main() {
//	ch := make(chan int)
//
//	// This goroutine will block forever
//	// because there are no receivers for the channel.
//	go func() {
//		ch <- 10
//		close(ch) // Closing the channel avoids deadlock
//	}()
//
//	// Sleep for a bit to allow the goroutine to start
//	time.Sleep(time.Second)
//
//	// Receive the data from the channel
//	for value := range ch {
//		fmt.Println("Received data:", value)
//	}
//	// Avoid sending on a closed channel to prevent panic
//	//ch <- 20
//}

//Objective: Implement a pipeline pattern where data flows
//through multiple stages of processing.

//func main() {
//	nums := gen(2, 3)
//
//	square := square(nums)
//
//	double := double(square)
//
//	for result := range double {
//		fmt.Println("Result:", result)
//	}
//}
//
//func gen(nums ...int) chan int {
//	out := make(chan int)
//	go func() {
//		for _, n := range nums {
//			out <- n
//		}
//		close(out)
//	}()
//	return out
//}
//
//func square(in chan int) chan int {
//	out := make(chan int)
//	go func() {
//		for n := range in {
//			out <- n * n
//		}
//		close(out)
//	}()
//	return out
//}
//
//func double(in chan int) chan int {
//	out := make(chan int)
//	go func() {
//		for n := range in {
//			out <- n * 2
//		}
//		close(out)
//	}()
//	return out
//}

//Objective: Learn how to handle closed channels properly in Go.

//func main() {
//	ch := make(chan int)
//
//	go func() {
//		defer close(ch)
//		for i := 0; i < 3; i++ {
//			ch <- i
//		}
//	}()
//
//	for value := range ch {
//		fmt.Println("Received data:", value)
//	}
//}

// Objective: Combine multiple channels with a timeout using the select statement.
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	go func() {
//		time.Sleep(1 * time.Second)
//		ch1 <- 1
//	}()
//
//	go func() {
//		time.Sleep(2 * time.Second)
//		ch2 <- 2
//	}()
//
//	for i := 0; i < 2; i++ {
//		select {
//		case val := <-ch1:
//			fmt.Println("Received from ch1:", val)
//		case val := <-ch2:
//			fmt.Println("Received from ch2:", val)
//		case <-time.After(3 * time.Second):
//			fmt.Println("Timeout!")
//		}
//	}
//}

// Objective: Implement a simple throttling mechanism using channels.

//func main() {
//	throttle := make(chan time.Time, 3)
//
//	for i := 0; i < 3; i++ {
//		throttle <- time.Now()
//	}
//
//	go func() {
//		for t := range time.Tick(1 * time.Second) {
//			throttle <- t
//		}
//	}()
//
//	for i := 0; i < 10; i++ {
//		<-throttle // Block until we can process the next task
//		fmt.Println("Processing task", i, "at", time.Now())
//	}
//
//}

// Objective: Implement a producer-consumer problem using channels.

func main() {
	tasks := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go consumer(i, tasks, &wg)
	}

	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()
}

func consumer(id int, tasks chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		fmt.Printf("Consumer %d processed task %d\n", id, task)
	}
}
