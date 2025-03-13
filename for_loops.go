package main

import "fmt"

func main(){
	// angka := 1

	// for angka <= 10{
	// 	fmt.Println("Perulangan ke ",angka)
	// 	angka++
	// }

	names := []string{"Farid", "Adhwa", "Putra", "Pratama"}
	// for i, name := range names{
	// 	fmt.Println("Index", i, "Nama", name)
	// }
	for _, name := range names {
		fmt.Println(name)
	}
	
}