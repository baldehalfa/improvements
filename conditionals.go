package main

import "fmt"

func main() {
	x := 10
	y := 10

	//  If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	// else if
	color := "green"

	if color == "green" {
		fmt.Println("color is green")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT green or blue")
	}

	// Switch
	switch color {
	case "green":
		fmt.Println("color is green")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT green or blue")
	}

}
