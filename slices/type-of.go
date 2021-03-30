package main

import (
	"fmt"
	"reflect"
)

func main()  {
	var names []string
	fmt.Println(reflect.TypeOf(names).Kind())
	{
		var names [5]string
		fmt.Println(reflect.TypeOf(names).Kind())
	}

}
