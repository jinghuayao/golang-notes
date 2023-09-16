package main

import (
	"fmt"
)


type person struct {
	firstName string
	lastName string
	contactInfo // equivalent to: contactInfo contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}



func main() {
	fmt.Println("hello construct")
	var xy person
	fmt.Printf("%+v", xy)

	xy.firstName = "Xiaoyan"
	xy.lastName = "Wang"
	xy.contactInfo = contactInfo{email: "xiaoyanwang@gmail.com", zipCode: 95132}

	fmt.Println("")
	xy.print()


	jy := person{
		firstName: "Jinghua",
		lastName: "Yao",
		contactInfo: contactInfo{
			email: "jinghuayao@gmail.com",
			zipCode: 95132, // this comma is needed
		}, // this comma is needed
	}
	fmt.Println("")
	jy.print()

	jy.updateFirstName("Chunliang")
	fmt.Println("")
	jy.print()
}



func (p person) print() {
	fmt.Printf("%+v", p)
}


func (p person) updateFirstName(*newFirstName string) {
	p.firstName = newFirstName
}