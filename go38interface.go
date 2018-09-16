package main

import (
	"fmt"
)

type goodPerson interface {
	work(hours int) string
}

type Emp struct {
	name, org string
	income    int
}

type Student struct {
	name, rollnum string
	age           int
}

func (s Student) work(hours int) string {
	return "Student working"
}

func (e Emp) work(hours int) string {
	return "Employee working"
}

func makePersonWorking(g goodPerson) {
	fmt.Printf("%T %s\n", g, g.work(10))
}

func main() {

	e := Emp{name: "abc", income: 1200}
	s := Student{name: "jack", age: 22}

	sl := make([]goodPerson, 0)
	sl = append(sl, e, s)

	for _, p := range sl {
		makePersonWorking(p)
	}

	fmt.Println(sl)

}
