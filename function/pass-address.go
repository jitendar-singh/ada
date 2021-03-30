package main

import "fmt"

func add(name *string, age *int) {

	*name = *name +" Singh"
	*age = *age +5
	//return
}
func main()  {
	name := "Jitendar"
	age := 25
	fmt.Println("before name was ",name," age was ",age)

	add(&name,&age)
	fmt.Println("after function name is ",name," age is ",age)

}