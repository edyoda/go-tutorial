package main

import (
	"fmt"
)

func main() {
	// var a int

	// println("Enter Num")
	// fmt.Scanf("%d", &a)

	var s string
	println("Enter String")
	fmt.Scanf("%s", &s)

	if (len(s) > 5) || (len(s) == 2) {
		println("Hello World", s)
	} else if len(s) == 5 {
		println("Great stuff")
	} else {
		println("Hello World Again", s)
	}

}
