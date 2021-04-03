package main

import "fmt"

func main() {
	fmt.Println(func(l int, b int) int {
		return l * b
	}(2, 4))
}
