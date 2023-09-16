package main

import (
	"fmt"
	"net/http"
	"os"
)



func main() {
	fmt.Println("")

	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	
	bs := make([]byte, 99999)

	resp.Body.Read(bs)
	fmt.Println(string(bs))
}