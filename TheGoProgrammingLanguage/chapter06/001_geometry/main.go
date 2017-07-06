package main

import (
	"fmt"

	"./geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}

	fmt.Println(geometry.Distance(p, q))
	fmt.Println(q.Distance(p))

	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())
}
