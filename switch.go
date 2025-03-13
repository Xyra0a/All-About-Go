package main

import (
	"fmt"
)

func main(){
	nama := "Budi bas"

	// switch nama{
	// case nama : 
	// 	fmt.Println("Halo", nama)
	// default : 
	// 	fmt.Println("Namanya siapa?")
	// }

	// short statement

	// switch length := len(nama); length > 3{
	// case true :
	// 	fmt.Println("Nama terlalu panjang")
	// default : 
	// 	fmt.Println("Mantap")
	// }

	length := len(nama)
	switch{
	case length <= 8:
		fmt.Println("Nama terlalu pendek")
	case  length > 16:
		fmt.Println("Nama terlalu panjang")
	default : 
		fmt.Println("Nama kosong")
	}
}