package main

import "fmt"

func main() {
	var a [5]int
	var hello [5]string
	newhello := [5]{"Hello","Hello","Hello","Hello"}
	
	a[0] = 9
	a[1] = 4
	a[2] = 6
	a[3] = 77
	a[4] = 9

	for idx := range hello {
		hello[idx] = "hello" + string('1'+idx)
	}

	for idx, v := range a {
		println(v)
		a[idx] *= 2
	}
	fmt.Println(hello)
}
