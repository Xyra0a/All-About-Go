package main 

import "fmt"


func getNamesCustomer()(firstName,lastName,fullName string){
	firstName  = "Tama"
	lastName = "Pratama"
	fullName = "Tama Pratama"

	return firstName,lastName,fullName
}
func main(){
	// Named Return Values
	fmt.Println(getNamesCustomer())
	// end Named Return Values
}