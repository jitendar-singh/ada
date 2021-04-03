package main

import "fmt"

func main() {
	var employee = make(map[string]int)
	employee["Jitendar"] = 1
	employee["Monalisa"] = 2
	employee["Gajanan"] = 3
	employee["Rathi"] = 4

	for key, value := range employee {
		fmt.Println("Deleting employee ", value)
		delete(employee, key)
	}
	for _, value := range employee {
		fmt.Println(value)
	}

}
