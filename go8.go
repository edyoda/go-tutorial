package main

import "fmt"

// func f() int {
// 	return 5
// }

func f() *int {
	v := 1
	return &v
}

func main() {
	a := 1
	p := &a
	fmt.Println(*p)
	*p = 22
	fmt.Println(a)
	c := &p
	fmt.Println(**c)

	var p1 = f()
	fmt.Println(*p1)
}
