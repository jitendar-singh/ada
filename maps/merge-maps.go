package main

import (
	"fmt"
)

func main() {
	map1 := map[string]int{"Jitendar": 1, "Sunny": 2, "Monalisa": 3}
	fmt.Println(map1)
	{
		func() {
			map2 := map[string]int{"Rekha": 22, "Sikha": 33, "Mika": 44}
			for key, value := range map2 {
				map1[key] = value
			}
			fmt.Println(map1)
		}()
	}

}
