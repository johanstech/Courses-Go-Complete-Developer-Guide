package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// person1 := person{"Bob", "Bobson"}
	person1 := person{
		firstName: "Bob",
		lastName:  "Bobson",
		contact: contactInfo{
			email:   "bob.bobson@mail.com",
			zipCode: 12345,
		},
	}
	person1.print()

	//? Updating values at pointers is so common that there's a shortcut:
	//? We can pass a value to a receiver function expecting a pointer
	//? as long as the value type and the pointer type are the same
	//? eg. person and *person
	person1.updateName("Bobbiest")
	person1.lastName = "Boblington"
	person1.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
