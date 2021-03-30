package main

import "fmt"

func main()  {
	i := 0
	for ;i<10;i++{
		fmt.Println(i+1)
	}

	for k:=0;k<10;k++{
		fmt.Println(k+1)
	}

	j := 0
	for j<10{
		fmt.Println(j+1)
		j++
	}

	q := 0
	for {
		if q >=10{
			break
		}
		fmt.Println(q + 1)
		q++
	}
}