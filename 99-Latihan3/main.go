package main

import (
	"fmt"
	"strconv"
)

func main() {

	var angka int

	fmt.Scan(&angka)

	angkaStr := strconv.Itoa(angka)

	if len(angkaStr) == 1 {
		angkaStr = "00" + angkaStr
	} else if len(angkaStr) == 2 {
		angkaStr = "0" + angkaStr
	}

	// Mengambil tiga digit pertama dari string
	d1 := angkaStr[:1]
	d2 := angkaStr[1:2]
	d3 := angkaStr[2:3]

	// tigaDigitPertama := angkaStr[:1] + " " + angkaStr[1:2] + " " + angkaStr[2:3]
	fmt.Printf("%s %s %s", d1, d2, d3)
}
