package main

import "fmt"

func main() {
	// Pointer adalah penunjuk data pertama di array para slice
	// Length adalah panjang dari slice, dan
	// Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity

	names := [...]string{"Farid", "Adhwa", "Putra", "Pratama", "Anjay", "Kelas", "a", "asd", "asdasd", "asda", "kasdas", "asdasd"}
	// Membuat slice dari array dimulai index low sampai index sebelum high
	slice := names[4:9]
	// 4 = index
	slice[0] = "Kontoll"

	fmt.Println(slice)
	fmt.Println(slice[0])
	fmt.Println(slice[1])

	// Membuat slide dari array dimulai index low sampai index akhir di array
	slice1 := names[5:]
	fmt.Println(slice1)

	// Membuat slice dari array dimulai index 0 sampai index sebelum high
	slice2 := names[:4]
	fmt.Println(slice2)

	// Membuat slice dari array dimulai index 0 sampai index akhir di array
	slice3 := names[:]
	fmt.Println(slice3)

	fmt.Println("=======================================================")
	fmt.Println("Ini buat function slice ")
	fmt.Println("=======================================================")


	funSlice1 := names[5:]
	funSlice1[1] = "Gacor"
	fmt.Println(append(funSlice1))

	funSlice2 := append(funSlice1, "Ini gacor baru")
	fmt.Println(funSlice2)
	// fmt.Println(len(funSlice1))
	// fmt.Println(cap(funSlice1))
}