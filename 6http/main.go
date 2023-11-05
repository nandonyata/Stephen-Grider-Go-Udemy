package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Err : ", err)
		os.Exit(1)
	}

	// 1
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	// 2
	// io.Copy(os.Stdout, resp.Body)

	// 3
	lw := logWriter{}
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {


	fmt.Println(string(bs))
	fmt.Println("JUST PROCESS THESE LENGTH OF DATA", len(bs))

	return len(bs), nil
}