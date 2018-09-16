package main

import (
	"fmt"
)

func main() {
	//var s string

	// for i := 1; i < len(os.Args); i++ {
	// 	s += os.Args[i]
	// }
	// fmt.Println(s)

	// for idx, arg := range os.Args {
	// 	fmt.Println(idx, arg)
	// }

	for idx, ch := range "hello world" {
		//fmt.Println(idx, ch)
		fmt.Printf("%d %T %c\n", idx, ch, ch)
	}
}
