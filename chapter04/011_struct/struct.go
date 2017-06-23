package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	DoB       time.Time
	Address   string
	Position  string
	Salary    int
	ManagerID int
}

func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

type Point struct{ X, Y int }

type address struct {
	hostname string
	port     int
}

func main() {
	var dilber Employee
	dilber.Salary -= 5000
	position := &dilber.Position
	*position = "Senior" + *position

	fmt.Println(dilber)

	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	hits := make(map[address]int)
	hits[address{"golong.org", 433}]++
}
