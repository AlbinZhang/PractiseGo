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

func main() {
	var dilber Employee
	dilber.Salary -= 5000
	position := &dilber.Position
	*position = "Senior" + *position

	fmt.Println(dilber)
}
