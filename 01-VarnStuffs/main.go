package main

import "fmt"

func main() {
	var a string
	var b int
	var c float32 // approx range 10^-101 to 10^90 (7 decimal digits)
	// var c float64  // approx range 10^-398 to 10^369 (16 decimal digits)
	var dd byte // char? (ASCII)
	var d rune  // char? (UTF-8)
	var e bool

	// const
	const cc float64 = 34.234523894345122

	a = "DuDu"
	b = 101
	c = 20.00034
	dd = 'A'
	d = '笑'
	e = true

	fmt.Printf("%s\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%f\n", c)

	fmt.Printf("%c\n", dd) // print the char
	fmt.Printf("%d\n", dd) // print the ASCII code

	fmt.Printf("%c\n", d) // print the char
	fmt.Printf("%d\n", d) // print the unicode code
	fmt.Printf("%U\n", d) // print the unicode

	fmt.Printf("%t\n", e)
	fmt.Printf("%.16f", cc)
}
