package main

import "fmt"

func main() {
	var twoDarray [4][4]int
	var count int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			count++
			twoDarray[i][j] = count
		}
	}

	fmt.Println(twoDarray)
}
