package main

import "fmt"

func incr(p *int) {
	*p++
}

func decr(p *int) {
	*p--
}

func main() {
	var counter int

	incr(&counter)
	incr(&counter)
	fmt.Println(counter)
	decr(&counter)
}
