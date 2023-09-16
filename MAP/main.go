/*
map
*/
package main

import (
	"fmt"
)


func main() {
	fmt.Println("hello, map")

	// // v1
	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf745", // the comma here is needed

	// }
	// fmt.Println(colors)
	
	// // v2
	// var colors map[string]string
	// fmt.Println(colors)

	// v3
	colors := make(map[string]string)
	fmt.Println(colors)


	colors["white"] = "#ffffff"
	colors["r"] = "red"
	colors["g"] = "green"
	colors["b"] = "blue"
	fmt.Println(colors)


	// delete key-val pairs 
	delete(colors, "white")
	fmt.Println(colors)



	freqs := make(map[int]int)
	freqs[1] = 10
	freqs[2] = 100
	freqs[3] = 1000
	freqs[3] = 0

	fmt.Println(freqs)
	delete(freqs, 1)
	fmt.Println(freqs)

	// iterate over map
	for key, val := range freqs {
		fmt.Printf("Key: %d, Value: %d\n", key, val)
	}

	for key, val := range colors {
		fmt.Printf("Key: %s, Value: %s\n", key, val)
	}

	printMap(freqs)
	fmt.Println("")

	// check membership
	if _, ok := freqs[3]; ok {
        // the key 3 exists within the map
        fmt.Println(freqs[3])
    }

}


func printMap(c map[int]int) {
	for key, val := range c {
		fmt.Printf("map-key: %d, map-value: %d", key, val)
	}
}


