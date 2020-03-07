package main

import "fmt"

func main() {

	emails := map[string]string{"Balde": "balde@gmail.com", "Mamadou": "mamadou@gmail.com"}

	emails["Alpha"] = "alpha@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Balde"])

	// Delete from map
	delete(emails, "Balde")
	fmt.Println(emails)
}
