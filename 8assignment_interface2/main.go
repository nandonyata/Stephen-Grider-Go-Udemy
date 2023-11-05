package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	nameFile := os.Args[1]
	openFIle, err := os.Open(nameFile)

	if err != nil {
		fmt.Println("ERR ; ", err)
	}

	io.Copy(os.Stdout, openFIle)
}