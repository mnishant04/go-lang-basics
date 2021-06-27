package main

import "fmt"

func main() {
	var name = "nishant"
	var nish = "aaj"
	var last = "mishra"
	job := "student"
	var age = 21
	first, middle := "nishant", "mishra"
	fmt.Println("name : ", name, last, "age = ", age, job, nish)
	fmt.Printf("%T %T %T %T %T \n", name, last, age, job, nish)
	fmt.Println(first, middle)

}
