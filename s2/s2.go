package main

import "fmt"

func main()  {
	var list = []int{10, 20, 30, 40,290,11,14,-15,-80}
	var inputNumber int
	fmt.Println("input number :")
	fmt.Scan(&inputNumber)

	listMap := make(map[int]int)

	for _,v:= range list{
		listMap[v]++
	}

	for _,v:=range list{
		if _, ok := listMap[inputNumber-v]; ok {
			fmt.Printf("found %d , %d \n", v, inputNumber-v)
			return
		}
	}
	fmt.Printf("not found")

}
