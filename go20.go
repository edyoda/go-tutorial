package main

import "fmt"

func main() {
	info := make(map[string]int)

	info["abc"] = 12
	info["def"] = 23

	fmt.Println(info)
	info["abc"]++
	delete(info, "abc")

	fmt.Println(info)

	res, error := info["abc"]
	fmt.Println(res, error)
}
