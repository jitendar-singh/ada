package main

import "fmt"

type employee struct {
	name string
	age int

	salary struct{
		ctc float64
		tax float64
		pf float64
		takehome float64
	}

}

func main()  {
	var emp employee
	emp.name = "Jitendar"
	emp.age = 32
	emp.salary.ctc = 1776000.00
	emp.salary.tax = (emp.salary.tax * 0.3)/12
	emp.salary.pf = (emp.salary.ctc * 0.5) * 0.12
	emp.salary.takehome = (emp.salary.ctc/12) - (emp.salary.tax + emp.salary.pf)

	fmt.Println(emp)
}
