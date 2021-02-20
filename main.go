package main

import "fmt"

// ANY OTHER TYPES that have a function with the same NAME AND OUTPUT of that function
// that are PASSED into an INTERFACEwill also be type BOT. In this case all
// Types that have a getGreeting func that return a string.
// Those types here are englishBot, and spanishBot
type bot interface {
	getGreeting() string
}

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
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

//THERE 2 functions are the same but with different type input.

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

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
