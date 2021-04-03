package main

import (
	"fmt"
	"reflect"
)

func main() {
	names := []string{
		"Sunny",
		"Runny",
		"Gunny",
		"Money",
	}
	fmt.Println(reflect.TypeOf(names).Kind())
	for index, value := range names {
		fmt.Println(index, ":", value)
	}
	for index := range names {
		fmt.Println(index)
	}
	for _, values := range names {
		fmt.Println(values)
	}
	for index := 0; index < len(names); index++ {
		fmt.Println(names[index])
	}
}
