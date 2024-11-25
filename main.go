package main

import (
	"fmt"
	"math/rand"
	"previewphase1/helper"
	"sync"
	"time"
)

func main() {
	// firstTask()
	// secondTask()
	thirdTask()
	// fourthTask()
	// fifthTask()
}

func firstTask() {
	go helper.PrintNumbers()
	go helper.PrintLetters()

	time.Sleep(3 * time.Second)
}

func secondTask() {
	var wg sync.WaitGroup

	if rand.Intn(2) == 0 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			helper.PrintNumbers()
		}()

		wg.Wait()

		wg.Add(1)

		go func() {
			defer wg.Done()
			helper.PrintLetters()
		}()

		wg.Wait()
	} else {
		wg.Add(1)

		go func() {
			defer wg.Done()
			helper.PrintLetters()
		}()

		wg.Wait()

		wg.Add(1)

		go func() {
			defer wg.Done()
			helper.PrintNumbers()
		}()

		wg.Wait()
	}
}

func thirdTask() {
	ch := make(chan int) // Unbuffered channel

	go helper.Produce(ch)
	go helper.Consume(ch)

	time.Sleep(300 * time.Millisecond)
}

func fourthTask() {
	ch := make(chan int, 5) // Buffered channel with a capacity of 5

	go helper.Produce(ch)
	go helper.Consume(ch)

	time.Sleep(300 * time.Millisecond)
}

func fifthTask() {
	oddNumber := make(chan int)
	evenNumber := make(chan int)
	isErr := make(chan error)

	go func() {
		for i := 1; i <= 25; i++ {
			if i > 20 {
				isErr <- fmt.Errorf("number %d is greater than 20", i)
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
