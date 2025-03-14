package main

import (
	"fmt"
)


func sumNumber(numbers ...int)int{
	total := 0

	for _, number := range numbers{
		total += number
	}

	return total
}
func main(){
	getNumber := sumNumber(10,10,10,10,10,10)
	fmt.Println(getNumber)

	getNumbers := []int{10,10,10,5}
	fmt.Println(sumNumber(getNumbers...))
}