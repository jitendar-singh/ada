package main

import "fmt"

func main() {
	var names = make([]string, 10)
	var lastNames = make([]string, 10, 20)

	fmt.Println(len(names), " ", cap(names))
	fmt.Println(len(lastNames), " ", cap(lastNames))
	fmt.Println(names)
	fmt.Println(lastNames)
}
