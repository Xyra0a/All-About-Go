package main 

import "fmt"

func main(){
	person := map[string]string{
		"nama":  "budi",
		"kelas": "12 RPL 1",
		"umur": "18",
	}
	person["nama"] = "Anjay"
	delete(person, "nama")

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["kelas"])
	fmt.Println(len(person))

	// city := map[string]string{
	// city := make(map[string]string){
	// 	"subdistrict" : "Pancoranmas",
	// 	"citizens_association" : "20",
	// }

	// fmt.Println(city["subdistrict"])

}