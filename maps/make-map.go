package main

import "fmt"

func main() {
	var employee = make(map[string]int)
	employee["Sunny"] = 1
	employee["Sony"] = 2

	fmt.Println(employee)

	student := make(map[string]int)
	student["Jitendar"] = 1
	student["Mona"] = 2

	fmt.Println(student)

}
