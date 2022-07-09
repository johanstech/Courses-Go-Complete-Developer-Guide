//? main is for executables, any other name is for reusable
package main

import "fmt"

//? executables always need a func main()
func main() {
	//* Long way of declaring variable
	// var message string = "Hello world!"

	//* Walrus operator := is only used when declaring a variable, not reassigning
	message := "Hello world!"

	fmt.Println(message)
}
