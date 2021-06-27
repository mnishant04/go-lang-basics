package main

import "fmt"

func main() {
	//map is a key value pair
	//define map
	student := make(map[string]string)

	//assigning value to map
	student["nishant"] = "900000000"
	student["pradeep"] = "900000001"
	student["sangam"] = "900000002"
	fmt.Println(student)           //complete map printing
	fmt.Println(student["nishant"]) //individual key pair printing
	fmt.Println(len(student))      //length of map 3

	//delete from map
	delete(student, "pradeep")
	fmt.Println(student)      //complete map printing
	fmt.Println(len(student)) //length of map 2


	//direct assign

	emails :=map[string]string{"nishant":"nishant@gmail.com", "pradeep":"pradeep@gmail.com", "sangam":"sangam@gmail.com"}
	fmt.Println(emails)
	emails["pratik"] = "pratik@gmail.com"
	fmt.Println(emails)

	// this can be used for varification
}
