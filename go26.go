package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func rever(s string) string {
	if len(s) == 0 {
		return ""
	}
	return string(s[len(s)-1]) + rever(s[:len(s)-1])
}

func main() {
	//res := fact(10)
	//fmt.Println(res)

	fmt.Println(rever("hello"))
}
