package main

import "fmt"

func main() {
	var jam, menit, detik int
	// jam = 2
	// menit = 30
	// c = 5

	fmt.Print("Masukkan jam: ")
	fmt.Scan(&jam)
	fmt.Print("Masukkan menit: ")
	fmt.Scan(&menit)
	detik = jam*3600 + menit*60

	fmt.Printf("%djam %dmenit = %d detik", jam, menit, detik)
}
