package main

import "fmt"

func main() {
	fmt.Println("Hello World :D") // printline ofc
	fmt.Print("Hello World :D\n") // print

	// ignore this
	var a string
	fmt.Scan(&a)

	fmt.Printf("%s\n", a) // printformat
}
