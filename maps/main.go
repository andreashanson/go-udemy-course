package main

import "fmt"

func main() {
	fmt.Println("This is maps")
	// Declare 1
	var colors map[string]string
	
	// Declare 2
	colors2 := make(map[string]string)

	// Declare 3
	colors3 := map[string]string{
		"red": "#ff0000",
		"green": "#745777",
		"white": "#FFFFFF",
	}

	printMap(colors)
	printMap(colors2)
	printMap(colors3)

	//colors["black"] = "#999999"
	// printMap(colors)
	
	colors2["black"] = "#999999"
	printMap(colors2)

	colors3["blue"] = "#444444"
	printMap(colors3)
	
	// This will delete the key "black from map colors2"
	delete(colors2, "black")
	fmt.Println(colors2)

	// Same as above but keys are of type int and values still strings.
	names := make(map[int]string)
	names[0] = "Anna"
	names[1] = "Andreas"
	fmt.Println(names)
	delete(names, 1)
	fmt.Println(names)

}

	func printMap(m map[string]string) {
		fmt.Println("Printing map of strings")
		for key, value := range m {
			fmt.Println("key: " + key)
			fmt.Println("value: " + value)
		}
	}
