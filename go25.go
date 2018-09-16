package main

import "fmt"

func fruitGen() func() string {
	fruits := []string{"apple", "mango", "banana", "orange", "grapes"}
	count := 0
	return func() string {
		fruit := fruits[count]
		count = (count + 1) % 4
		return fruit
	}
}

func seqGen() func() int {
	count := 1
	return func() int {
		count++
		return count
	}
}

func main() {

	seqCtr := seqGen()

	fmt.Println(seqCtr())
	fmt.Println(seqCtr())
	fmt.Println(seqCtr())
	fmt.Println(seqCtr())

	seqCtr1 := fruitGen()

	fmt.Println(seqCtr1())
	fmt.Println(seqCtr1())
	fmt.Println(seqCtr1())
	fmt.Println(seqCtr1())
	fmt.Println(seqCtr1())
	fmt.Println(seqCtr1())
}
