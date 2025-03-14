package main

import (
	"fmt"
)


func addSum(numbers ...int)int{
	total := 0
	for _, number := range numbers{
		total += number
	}

	return total
}


func main(){
	fmt.Println(addSum(1, 2, 3, 4, 5))
}