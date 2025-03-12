package main

import "fmt"

func main() {
	var nama1 string = "Tama"
	var nama2 string = "Tama"

	var result bool = nama1 != nama2
	fmt.Println(result)
	var a int = 10
	var b int = 5
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
}
