package main

import (
	"fmt"
)

type Salary struct {
	Basic, HRA, TA float64
}
type Employee struct {
	FirstName, LastName, Email string
	Age int
	MonthlySalary []Salary
}

func main()  {
	e := Employee{
		FirstName: "Jitendar",
		LastName:  "Singh",
		Email:     "jitsingh@redhat.com",
		Age:       32,
		MonthlySalary: []Salary{
			{
				Basic: 45000,
				HRA:   12000,
				TA:    6000,
			},
		},
	}
	fmt.Println(e)
	for _,values := range e.MonthlySalary{
		fmt.Println(values)
	}
}
