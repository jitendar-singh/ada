package main

import "fmt"

func main()  {
	slice1 := []int{1,2,3,4,5}
	slice2 := []int{1,2,3,4}
	slice2 = append(slice2,slice1[len(slice1)-1:]...)
	fmt.Println(slice2)

}
