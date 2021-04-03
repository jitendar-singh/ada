package main

import "fmt"

type employee struct {
	name       string
	age        int
	department string
	yoe        int
}

func main() {
	var emp1 = employee{"Jitendar", 32, "QE", 6}
	fmt.Println(emp1)

	var emp2 = employee{
		name:       "Monalisa",
		age:        29,
		department: "Info Sec",
		yoe:        6,
	}
	fmt.Println(emp2)

	emp3 := employee{"Sandip", 35, "Private", 12}
	fmt.Println(emp3)
}
