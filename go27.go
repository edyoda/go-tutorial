package main

import "fmt"

func myfunc1() {
	fmt.Println("More here")
}

func myfunc() {
	fmt.Println("Finished Now")
}

func main() {
	defer myfunc()

	//defer myfunc1()

	fmt.Println("Is it finished")

}
