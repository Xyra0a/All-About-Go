package main

import "fmt"


// func helloWorld(){
// 	fmt.Println("Hello world")
// }

// function as paramater
// func helloWorld(firstName string, lastName string){
// 	fmt.Println("Hallo", firstName, lastName)
// }
// end function as paramater

// function return value
	// func helloWorld(firstName string) string{
	// 	return "Hallo "+firstName
	// }
// end function return value

// function return multiple values
	// func helloWorld(username,password string)(bool,string){
	// 	if username == "admin123" && password == "123" {
	// 		return true, "Login berhasil"
	// 	}else{
	// 		return false, "Login gagal"
	// 	}
	// }
	/*
		
	*/
	// func helloWorld()(string, string){
	// 	return "Tama", "Pratama"
	// }

// end function return multiple values



func main(){
	success, _ := helloWorld()
	fmt.Println(success)
	// fmt.Println(helloWorld("Tama"))
}