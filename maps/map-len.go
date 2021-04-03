package main

import "fmt"

func main() {
	students := make(map[string]int)
	students["Jitendar"] = 1
	students["Monalisa"] = 2
	students["Gada"] = 3
	fmt.Println(len(students))

	var employee map[string]int
	fmt.Println(len(employee))

	janitor := map[string]int{"sashi": 1, "Rama": 2}
	fmt.Println(janitor)
}
