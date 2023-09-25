package main

import (
	"fmt"
)

func selectionSortAsc(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		temp := arr[i]
		arr[i] = arr[min]
		arr[min] = temp
	}
}

func print(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func main() {
	arr := make([]int, 3)

	fmt.Print("Input data 1: ")
	fmt.Scan(&arr[0])
	fmt.Print("Input data 2: ")
	fmt.Scan(&arr[1])
	fmt.Print("Input data 3: ")
	fmt.Scan(&arr[2])

	fmt.Println("Original Array:")
	print(arr)

	fmt.Println("Ascending Sorted Array:")
	selectionSortAsc(arr)
	print(arr)
}
