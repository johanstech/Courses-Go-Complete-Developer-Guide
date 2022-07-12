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

	person1Pointer := &person1
	person1Pointer.updateName("Bobbiest")
	person1.lastName = "Boblington"
	person1.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
