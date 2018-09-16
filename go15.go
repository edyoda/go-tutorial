package main

import "fmt"

func main() {
	s := make([]string, 3)
	// s[0] = "hello"
	// s[1] = "world"
	// s[2] = "here"

	s = append(s, "I")
	s = append(s, "come", "again")

	fmt.Println(s)
}
