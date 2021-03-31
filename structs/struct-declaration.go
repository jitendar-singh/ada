package main

import "fmt"
type rectangle struct {
	lenght float64
	breadth float64
	height float64
	color string
}

func main(){

	fmt.Println(rectangle{
		lenght:  1,
		breadth: 22.5,
		height:  3,
		color: "red",
	})
}
