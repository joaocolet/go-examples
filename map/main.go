package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	//var colors map[string]string
	//colors := make(map[string]string)

	//colors["red"] = "#ff0000"
	//colors["green"] = "#00ff00"
	//colors["blue"] = "#0000ff"
	//fmt.Println(colors)

	//delete(colors, "red")
	//fmt.Println(colors)
	//colors["red"] = "#ff0000"
	//fmt.Println(colors)

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
