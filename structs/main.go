/*
pointers in Golang

取memory地址:
&variable: give me the memory address of the value this variable is pointing at
取memory地址对应的值：
*pointer: give me the value this memory address is pointing at


Value Types VS Reference Types

Value Types: int, float, string, bool, struct
Use pointers to change their things in a function

Reference Types: slices, maps, channels, pointers, functions
Do not worry about pointers with these types.

*/

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
	fmt.Println("")
	
	// Go pointer shortcut
	jy.updateFirstName("Chunliang")
	fmt.Println("")
	jy.print()


	jyPointer := &jy
	jyPointer.updateFirstName("XXX")
	jy.print()
	jy.updateFirstName("YYY")
	fmt.Println("")
	fmt.Printf("%+v", jy)
}


func (p person) print() {
	fmt.Println()
	fmt.Printf("%+v", p)
}

/*
Go is a pass-by-value language

*type: type description, meaning working with pointer to type instance

*memory_address; &value
*variable: give me the value of the memory address the variable is pointing to
&variable: give me the memory address (in RAM) for the value the variable is pointing to
*/
func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}