package main

import "fmt"

func main() {
	
	var age int32 = 17
	const isCool = true
	var size float32 = 2.3

	name, email := "Alpha", "alfa@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
