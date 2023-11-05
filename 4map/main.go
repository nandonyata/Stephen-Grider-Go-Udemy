package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":  "#ff000",
	// 	"blue": "#blseas313",
	// }

	// var colors map[string]string
	// OR
	colors := make(map[string]string)

	// add or update key
	colors["black"] = "#se1314"
	colors["white"] = "#ffff"

	// delete a key
	delete(colors, "black")

	// fmt.Println(colors)
	print(colors)
}

func print(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for", color + "is", hex)
	}
}