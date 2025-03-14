package main

import "fmt"

func getHello(name string, filter func(string)string){
		fmt.Println(filter(name))
}

func getDataFromHello(name string)string{
	if name == "Anjing" {
		return "Nama anda mengandung unsur kekerasan"
	}else{
		return name
	}
}

func main(){
	getHello("Tama",getDataFromHello) 

	filter := getDataFromHello
	getHello("Anjing",filter)
}