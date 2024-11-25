package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// firstTask()
	// secondTask()
	// thirdTask()
	// fourthTask()
	fifthTask()
}

func firstTask() {
	go printNumbers()
	go printLetters()

	time.Sleep(3 * time.Second)
}

func secondTask() {
	var wg sync.WaitGroup

	if rand.Intn(2) == 0 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			printNumbers()
		}()

		wg.Wait()

		wg.Add(1)

		go func() {
			defer wg.Done()
			printLetters()
		}()

		wg.Wait()
	} else {
		wg.Add(1)

		go func() {
			defer wg.Done()
			printLetters()
		}()

		wg.Wait()

		wg.Add(1)

		go func() {
			defer wg.Done()
			printNumbers()
		}()

		wg.Wait()
	}
}

func thirdTask() {
	c := make(chan int) // Unbuffered channel

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i // This will block until the receiver reads from the channel
		}

		close(c)
	}()

	for row := range c { // Receiver will block until data is available
		fmt.Println(row)
	}
}

func fourthTask() {
	c := make(chan int, 5) // Buffered channel with a capacity of 5

	go func() {
		for i := 1; i <= 10; i++ {
			c <- i // This will not block until the buffer is full
		}

		close(c)
	}()

	for row := range c { // Receiver will block if buffer is empty
		fmt.Println(row)
	}
}

func fifthTask() {
	oddNumber := make(chan int)
	evenNumber := make(chan int)
	isErr := make(chan error)

	go func() {
		for i := 1; i <= 25; i++ {
			if i > 20 {
				isErr <- fmt.Errorf("Number %d is greater than 20\n", i)
				continue
			}

			if i%2 == 0 {
				evenNumber <- i
			} else {
				oddNumber <- i
			}
		}

		close(oddNumber)
		close(evenNumber)
	}()

	for i := 0; i < 25; i++ {
		select {
		case odd := <-oddNumber:
			fmt.Println("Receive an odd number:", odd)
		case even := <-evenNumber:
			fmt.Println("Receive an even number:", even)
		case isErr := <-isErr:
			fmt.Printf("Error: %s", isErr)
		}

	}
}

func printNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func printLetters() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for _, row := range letters {
		fmt.Println(row)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}
