package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex = &sync.Mutex{}
	//mapmy := make(map[string]int)
	var counter int

	//mapmy["0"] = 0
	go func() {
		for i := 0; i < 100000; i++ {
			mutex.Lock()
			counter = counter + 1
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 100000; i++ {
			mutex.Lock()
			counter = counter - 1
			mutex.Unlock()
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println(counter)

}
