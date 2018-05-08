package main

import (
	"time"
	"fmt"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee

	dilbert.Salary = 5000
	fmt.Println(dilbert.Salary)

	position := &dilbert.Position
	*position = "Senior developer"
	fmt.Println(dilbert.Position)

	newEmployee := EmployeeByID(1)
	fmt.Println(newEmployee)

}

func EmployeeByID(id int) *Employee {
	return &Employee{ID: id, Name: "John"}
}
