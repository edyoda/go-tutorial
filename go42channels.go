package main

import "fmt"

func main() {
	numgenerator := make(chan int)
	numbercuber := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			numgenerator <- i
		}
		//close(numgenerator)
	}()

	go func() {
		for {
			i := <-numgenerator
			numbercuber <- i * i * i
		}
		//close(numbercuber)
	}()

	for x, err := range numbercuber {
		fmt.Println(x, err)
	}

}
