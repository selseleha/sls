package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str1 string
	var str2 string

	fmt.Println("input string 1 :")
	fmt.Scan(&str1)
	fmt.Println("input string 2 :")
	fmt.Scan(&str2)
	input1 := make(map[int32]int)
	input2 := make(map[int32]int)

	for _,v:= range str1 {
		input1[v]++
	}

	for _,v:= range str2 {
		input2[v]++
	}


	eq := reflect.DeepEqual(input1, input2)
	if eq {
		fmt.Println("They're equal.")
	} else {
		fmt.Println("They're unequal.")
	}
}




