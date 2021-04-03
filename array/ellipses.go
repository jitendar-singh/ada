package main

import "fmt"

func main() {
	names := [...]string{
		"Jitendar",
		"Sandip",
		"Kuldip",
		"Surjit",
	}
	for index, value := range names {
		fmt.Println("names[", index, "]=", value)
	}
}
