package main

import "fmt"

func main() {

	var employee = map[string]int{
		"Jitendar": 1,
		"Monalisa": 2,
		"Dudhi":    3,
		"Sudhi":    4,
	}
	for _, value := range employee {
		fmt.Println(value)
	}
	delete(employee, "Dudhi")
	for _, value := range employee {
		fmt.Println(value)
	}

}
