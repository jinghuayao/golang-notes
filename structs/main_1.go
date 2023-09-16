/*
struct and ways of creation
*/

package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}



func main() {
	fmt.Println("hello struct")

	// option 1 depends on the order of fields; error-prone
	alex := person{"Alex", "Anderson"}
	fmt.Println(alex)
	// option 2
	// key:value pairs, order does not matter
	xy := person{firstName: "Xiaoyi", lastName: "Zhang"}
	fmt.Println(xy)

	jy := person{lastName: "Yao", firstName: "Jinghua"}
	fmt.Println(jy)

	// option 3
	var xw person
	fmt.Println(xw)
	fmt.Printf("%+v", xw)
	fmt.Println("---")
	xw.firstName = "Xiaoyan"
	xw.lastName = "Wang"
	fmt.Println(xw)

	fmt.Printf("%+v", xw)

	xw.firstName = "XY"
	xw.lastName = "W"
	fmt.Println("")
	fmt.Printf("%+v", xw)

}