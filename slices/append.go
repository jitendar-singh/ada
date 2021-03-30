package main

import "fmt"

func main()  {
	var names = make([]string,5)
	names[0]="Jitendar"
	names =  append(names,"Sony","Jitu")
	fmt.Println(names)
	fmt.Println(len(names))
}