package main

import "fmt"
import "os"
import "io"

func main() {

	a := 1
	b := 2

	defer fmt.Println(a + b)

	a = 4
	defer fmt.Println(a + b)

	f, err := os.Open("go25.go")
	if err != nil {
		return
	}
	defer f.Close()

	dst, err := os.Create("text")
	io.Copy(dst, f)

}
