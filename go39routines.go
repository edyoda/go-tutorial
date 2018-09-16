package main

import (
	"fmt"
	"time"
)

func rotate(delay time.Duration) {

	for {
		for _, r := range "-\\|/" {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}

}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {

	go rotate(100 * time.Millisecond)
	f := fib(50)
	fmt.Println(f)

}
