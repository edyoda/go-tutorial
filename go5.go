package main

import "fmt"
import "math"

const s string = "hello constant"

func main() {
	fmt.Println(s)
	const n = 10000
	fmt.Println(math.Cos(n))
	var i int
	i = 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	num1 := 10
	num2 := 3

	if num1%num2 == 0 {
		fmt.Println("Div")
	} else if num1 == 10 {
		fmt.Printf("Ten")
	} else {
		fmt.Println("Else")
	}

	for {

	}

	for j := 1; j <= 10; j++ {

	}

}
