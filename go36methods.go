package main

import "fmt"

type traiangle struct {
	base, height float32
}

func (t traiangle) area() float32 {

	area := .5 * t.base * t.height
	t.base *= 2
	t.height *= 3
	return area
}

func (t *traiangle) areaptr() float32 {
	area := .5 * t.base * t.height
	t.base *= 2
	t.height *= 3
	return area
}

func (t traiangle) getter() (float32, float32) {
	return t.base, t.height
}

func main() {

	t := traiangle{base: 10, height: 20}

	fmt.Println(t.area())

	fmt.Println(t)

	tp := &t

	fmt.Println(tp.areaptr())

	fmt.Println(t.area())

	triangles := make([]traiangle, 0)
	triangles = append(triangles, t)

	for idx, traiangle := range triangles {

	}

}
