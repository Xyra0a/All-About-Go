package main 

import "fmt"

func main(){
	// var nama[3] string
	// var nama[4] string
	// nama[0] = "Farid"
	// nama[1] = "Adhwa"
	// nama[2] = "Pratama"
	// nama[3] = "Putra"

	// fmt.Println(len(nama[0]))
	// fmt.Println(array(nama[1]))
	// fmt.Println(nama[2])
	// fmt.Println(nama[3])

	var values = [...] int{ // ... berguna jika mau menyimpan tanpa ditentukan muatannya berapa
		10,
		20,
		30,
		80,
	}
	// var values = [3] int{
	// 	10,
	// 	20,
	// 	30,
	// }
	fmt.Println(len(values))
	fmt.Println(values[2])
	values[1] = 50 //mengubah data array menggunakan index
	fmt.Println(values)




}