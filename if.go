package main

import "fmt"

func main(){

	nama := "Tama Bagas"

	// if(nama == "Tama"){
	// 	fmt.Println("Halo", nama)
	// }else{
	// 	fmt.Println("Namanya tidak ada")
	// }

	if length := len(nama); length < 8{
		fmt.Println("Nama terlalu pendek")
	}else{
		fmt.Println("Nama di approved")
	}
}