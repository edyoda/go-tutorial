package main

import "fmt"

type Slice []string

func main() {
	info1 := make(map[string]int)

	info1["abc"] = 12
	info1["def"] = 23
	info1["ghi"] = 23

	info := make(map[int]Slice)

	for name, age := range info1 {
		info[age] = append(info[age], name)
	}

	fmt.Println(info)

}
