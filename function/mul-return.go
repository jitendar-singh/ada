package main

import "fmt"

func rectangle(len int,breadth int)(area int,perimeter int)  {

	area = len * breadth
	perimeter = 2*(len+breadth)
	return
}
func main()  {
	area,perimeter := rectangle(10,20)
	fmt.Println(area,perimeter)
}