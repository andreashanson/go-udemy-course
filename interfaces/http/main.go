package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
)

func main() {
	fmt.Println("Interfaces with http")
	r, err := http.Get("https://www.example.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	// Way one
	//bs := make([]byte, 99999)
	//r.Body.Read(bs)
	//fmt.Println(string(bs))

	// Way two
	io.Copy(os.Stdout, r.Body)


}