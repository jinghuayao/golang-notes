/*
Nested struct
*/
package main

import (
	"fmt"
)

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}



func main() {
	fmt.Println("hello construct")
	var xy person
	fmt.Printf("%+v", xy)

	xy.firstName = "Xiaoyan"
	xy.lastName = "Wang"
	xy.contact = contactInfo{email: "xiaoyanwang@gmail.com", zipCode: 95132}

	fmt.Println("")
	fmt.Printf("%+v", xy)


	jy := person{
		firstName: "Jinghua",
		lastName: "Yao",
		contact: contactInfo{
			email: "jinghuayao@gmail.com",
			zipCode: 95132, // this comma is needed
		}, // this comma is needed
	}
	fmt.Println("")
	fmt.Printf("%+v", jy)
}