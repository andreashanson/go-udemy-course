package main

import ("fmt")

type shape interface {
	getArea() float64
}

type square struct{
	sideLength float64
}

type triangle struct{
	side float64
	base float64
}

func main() {
	fmt.Println("Interface test")
	t := triangle{side: 5, base: 8}
	s := square{sideLength: 4}

	printArea(s)
	printArea(t)

}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.side * t.base
}


