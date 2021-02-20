package main

import "fmt"

func main() {

	//Different ways to create a MAP

	// #1 create a map
	// var colors map[string]string

	//#2 Create a map
	// colors := make(map[string]string)

	// #3 Create a map
	colors := map[string]string{
		"red":   "ff0000",
		"green": "75b271",
		"white": "ffffff",
	}

	delete(colors, "white")

	fmt.Println(colors)
}
