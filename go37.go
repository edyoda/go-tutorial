package main

import "fmt"
import "strconv"

type Employee struct {
	age, salary        int
	name, organization string
	workstatus         bool
}

func (e *Employee) work() {
	e.workstatus = true
}

func createEmployee() Employee {
	e := Employee{age: 30, salary: 10000, name: "Raghav"}
	return e
}

func main() {

	emprecords := make(map[string]Employee)
	for i := 1; i < 10; i++ {
		emprecords["IBM00"+strconv.Itoa(i)] = createEmployee()
	}

	//fmt.Println(emprecords)

	var e Employee
	var ep *Employee

	for key, _ := range emprecords {
		e = emprecords[key]
		ep = &e
		ep.work()
		emprecords[key] = *ep
	}

	g := emprecords["IBM001"]

	g.workstatus = true
	emprecords["IBM001"] = g

	fmt.Println(emprecords)
}
