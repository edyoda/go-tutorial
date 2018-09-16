package main

import (
	"fmt"
	"os"
)

func myfunc() {
	fmt.Println("here")
}

func main() {

	_, err := os.Open("go28.god")
	fmt.Printf("%T", err)
	if err != nil {
		//fmt.Println(err)
		//defer myfunc1()
		defer myfunc()
		panic("****Ends Here****")
		myfunc()
	}

}
