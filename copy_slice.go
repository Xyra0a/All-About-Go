package main

import "fmt"

func main(){
	names := [...]string{"Farid", "Adhwa", "Putra", "Pratama", "Anjay", "Kelas", "a", "asd", "asdasd", "asda", "kasdas", "asdasd"}

	fromSlice := names[:]
    toSlice := make([]string, len(fromSlice), cap(fromSlice))
	
	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	fmt.Println(fromSlice)
}