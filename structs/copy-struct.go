package main

import "fmt"

type rectangle struct {
	lenght float64
	breadth float64
	height float64
}

func main()  {
	rec := rectangle{
		lenght: 12.4,
		breadth: 12,
		height: 12,
	}
	fmt.Println("Rec:",rec)
	rec2 := &rec
	fmt.Println("Rec2:",rec2)
	rec2.height = 14
	rec.lenght = 0
	fmt.Println("Rec:",rec)
	fmt.Println("Rec2:",rec2)

}
