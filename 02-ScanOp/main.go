package main

import "fmt"

func main() {
	var a, b int
	var c bool

	fmt.Print("Masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b : ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c (Bool) : ")
	fmt.Scan(&c)

	fmt.Printf("a + b = %d\n", a+b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %.3f\n", float32(a)/float32(b))
	fmt.Printf("a %% b = %d\n", a%b)
	fmt.Printf("!c = %t\n", !c)
}
