package main

import "fmt"

func main(){
	type nis string 

	var n nis = "12345"

	var warga string = "warga"
	var wargaFix nis = nis(warga)

	fmt.Println(wargaFix)
	fmt.Println(n)
}