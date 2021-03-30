package main

import "fmt"

func linearSearch(dataList []int,key int)bool{
	for _,item := range dataList{
		if item == key{
			return true
		}
	}
	return false
}

func main()  {
	items := []int{12,34,55,67,88,99,43,2,3,45,67,89,9}
	fmt.Print(linearSearch(items,0))
}