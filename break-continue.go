package main

import "fmt"

func main(){

	for i:= 0; i <=5; i++{
		// fmt.Println("Loopingan ke",i)
		if(i%5 == 0){
			continue
		}

		fmt.Println("Ini loopingan ke",i)
	}
}