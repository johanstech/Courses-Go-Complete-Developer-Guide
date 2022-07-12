package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"
	// colors["black"] = "#000000"

	// delete(colors, "white")

	//? Maps are basically key-value pairs
	//? All keys must be of same type, and all values must be of same type
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Hex code for", k, "is", v)
	}
}
