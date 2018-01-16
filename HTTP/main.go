package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)
	//create an empty byte slice with space for 99999 elements
	//the read function will not resize a slice if it is full, hence the massive slice
	resp.Body.Read(bs)
	//take byte slice, take the html contained within the resp.Body, and read that data into the byte slice (bs)
	fmt.Println(string(bs))
	//turn bs into string and print it.
}
