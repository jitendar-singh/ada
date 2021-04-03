package main

import "fmt"

type salary struct {
	basic, hra, ta float64
}

type employee struct {
	name          string
	age           int
	email         string
	monthlysalary []salary
}

func (e employee) EmpInfo() string {
	fmt.Println(e.name)
	fmt.Println(e.age)
	fmt.Println(e.email)

	for _, value := range e.monthlysalary {
		fmt.Println(value.basic)
		fmt.Println(value.hra)
		fmt.Println(value.ta)
	}
	return ""
}
func main() {
	e := employee{
		name:  "JItendar",
		age:   32,
		email: "sunnyconcise@gmail.com",
		monthlysalary: []salary{
			{
				basic: 23000.45,
				hra:   4350,
				ta:    786.3,
			}, {
				basic: 45000,
				hra:   5710,
				ta:    980,
			},
		},
	}
	fmt.Println(e.EmpInfo())
}
