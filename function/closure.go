package main

import "fmt"

func main()  {
	l := 10
	b := 20

	func(){
		var area int
		area = l *b
		fmt.Println(area)
	}()
}
