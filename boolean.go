package main

import "fmt"

func main() {
	var nilai1 int = 5
	var nilai2 int = 10
	var absen = 20

	var lulus = (nilai1 + nilai2) > 20
	var absensi = absen > 10

	var result = lulus || absensi
	fmt.Println(result)
}
