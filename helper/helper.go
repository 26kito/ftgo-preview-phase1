package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func PrintLetters() {
	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	for _, row := range letters {
		fmt.Println(row)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
	}
}

func Produce(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i // This will block until the receiver reads from the channel
		time.Sleep(3 * time.Millisecond)
	}

	close(ch)
}

func Consume(ch chan int) {
	for row := range ch { // Receiver will block until data is available
		fmt.Println(row)
	}
}
