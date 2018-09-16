package main

import (
	"fmt"
	"time"
)

func main() {
	i := "2"

	switch i {
	case "1":
		fmt.Println(`One`)
	case "2":
		fmt.Println("Two")
	default:
		fmt.Println("Default")
	}

	my_time := time.Now().Weekday()
	fmt.Println(my_time)

}
