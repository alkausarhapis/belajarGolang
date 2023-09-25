package main

import "fmt"

func main() {
	// Soal A slide 30
	fmt.Println("==== SOAL A ====")
	var x, y float64

	fmt.Print("Masukkan Nilai x : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan Nilai y : ")
	fmt.Scan(&y)

	// x = 5
	// y = 46
	formula := 1/(3*x*x+10) + 10*y + 7

	fmt.Printf("%.14f\n", formula)

	// Soal B slide 31
	fmt.Println("\n==== SOAL B ====")
	fmt.Print("Masukkan Nilai x : ")
	fmt.Scan(&x)
	formulaB := (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)
	if x != 0 {
		fmt.Printf("%.16f", formulaB)
	} else {
		fmt.Printf("%d", int(formulaB))
	}
}
