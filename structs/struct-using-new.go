package main

import "fmt"

type student struct {
	name string
	age int
	stream string
}

func main()  {
	var std1 = new(student)
	std1.name = "Jitendar"
	std1.age = 32
	std1.stream = "Comp Sc."

	fmt.Println(*std1)
}
