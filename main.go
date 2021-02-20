package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {

	//Different ways to create a MAP

	// #1 create a map
	// var colors map[string]string

	//#2 Create a map
	// colors := make(map[string]string)

	// #3 Create a map
	// 	colors := map[string]string{
	// 		"red":   "ff0000",
	// 		"green": "75b271",
	// 		"white": "ffffff",
	// 	}

	// 	printMap(colors)
	// }

	// func printMap(c map[string]string) {
	// 	for color, hex := range c {
	// 		fmt.Println("Hex code for", color, "is", hex)
	// 	}

	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printGreeting(sb)

}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting
	return ("Hi, There!")
}

// If you are not going to use the Placeholder  you can remove it OPTIONAL
func (spanishBot) getGreeting() string {
	return ("Hola!")
}
