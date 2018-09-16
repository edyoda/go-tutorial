package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary float32
}

func main() {
	e := Employee{"abc", 35, 10000}
	fmt.Println(e)

	f := Employee{name: "def", age: 33, salary: 10000}
	fmt.Println(f)

	ep := &e
	ep.age += 1
	fmt.Println(ep)

	emps := make([]Employee,0)
	emps = append(emps,e,f)

	fmt.Println(emps)
}
