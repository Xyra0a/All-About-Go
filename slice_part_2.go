package main

import "fmt"

func main(){
	// names := [...]string{"Farid", "Adhwa", "Putra", "Pratama", "Anjay"}
	// slice := names[3:5]

	// fmt.Println(slice)

	var newSlice []string = make([]string, 2,5)
	newSlice[0] = "Gacor"
	newSlice[1] = "Gacor_kali"
	
	// newSlice[5] = "Gacor kang banget"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	var slice2 []string = append(newSlice, "Eko")
	slice2[1] = "Anjay"

	fmt.Println(slice2)
	fmt.Println(newSlice)

	fmt.Println(slice2[1])
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))


}