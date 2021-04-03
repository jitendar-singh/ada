package main

import (
	"fmt"
	"reflect"
)

func main() {

	var names [2]string
	var age [5]int

	fmt.Println(reflect.TypeOf(age).Kind())
	fmt.Println(reflect.TypeOf(names).Kind())

}
