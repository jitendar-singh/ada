package main

import (
	"fmt"
	"reflect"
)

func main() {
	var employee = map[string]int{}
	fmt.Println(reflect.TypeOf(employee).Kind())
}
