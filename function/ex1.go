package main

import "fmt"

func sum(a int, b int) (res int) {
	res = a + b
	return res
}
func main() {
	var x, y int
	fmt.Println("Enter two numbers")
	fmt.Scanf("%d %d", &x, &y)
	fmt.Println("sum of ", x, " and ", y, " is ", sum(x, y))
}
