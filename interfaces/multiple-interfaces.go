package main

import "fmt"

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSides()
}
type Pentagon int

func (p Pentagon) Perimeter()  {
	fmt.Println("Perimeter of Pentagon:",5*p)
}
func (p Pentagon) NumberOfSides()  {
	fmt.Println("Pentagon has 5 Sides")
}

func main()  {

}