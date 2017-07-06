package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	//w.Circle.Point.X = 8
	w.Y = 8
	//w.Circle.Point.Y = 8
	w.Radius = 5
	//w.Circle.Radius = 5
	w.Spokes = 20

	ww := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Println(ww)
}
