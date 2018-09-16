package main

import "fmt"

func main() {
	fruits := []string{"mango", "apple", "pineapple"}

	morefruits := [5]string{}
	morefruits[0] = "orange"

	fmt.Println(morefruits)
	fmt.Println(fruits)

	morefruits[6] = "jhasjash"
	fmt.Println(morefruits)
}
