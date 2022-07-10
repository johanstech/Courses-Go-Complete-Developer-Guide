package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// person1 := person{"Bob", "Bobson"}
	person1 := person{firstName: "Bob", lastName: "Bobson"}
	fmt.Println(person1)

	person1.firstName = "Bobbiest"
	person1.lastName = "Boblington"
	fmt.Printf("%+v", person1)
}
