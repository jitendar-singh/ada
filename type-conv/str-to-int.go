package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main()  {
	strVar := "100"
	intvar, _ := strconv.Atoi(strVar)

	strVar1 := "-52541"
	intVar1, _ := strconv.ParseInt(strVar1,10,32)

	strVar2 := "101010101010101010"
	intVar2, _ := strconv.ParseInt(strVar2, 10, 64)
	fmt.Println(reflect.TypeOf(intvar).String())
	fmt.Println(reflect.TypeOf(intVar1).String())
	fmt.Println(reflect.TypeOf(intVar2).String())

}
