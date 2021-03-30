package main

import "fmt"

func main()  {
	names := [...]string{
		"Sunny",
		"Sony",
		"Nono",
	}
	family := names
	names[0] = "Jitendar"
	fmt.Println(names)
	fmt.Println(family)
}
