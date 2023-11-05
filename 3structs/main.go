package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	   contactInfo
}

func main() {
	// nando := person{
	// 	firstName: "Nando",
	// 	lastName:  "Nyata Pasti",
	// }

	// nando.lastName = "pewpew"

	// fmt.Println(nando)
	// fmt.Printf("%+v", nando)

	alex := person{
		firstName: "alex",
		lastName:  "bernabo",
		contactInfo: contactInfo{
			email:   "alex@mail.com",
			zipCode: 111},
	}


	// alexPointer := &alex
	// alexPointer.changeFirstName("nando")
	alex.changeFirstName("nando")

	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) changeFirstName(name string) {
	(*pointerToPerson).firstName = name
}