package main

import (
	"fmt"
)

func main() {
	size := readSize()
	array := readArray(size)
	printReversed(array)
}

func readSize() int {
	var n int
	fmt.Scanf("%d\n", &n)
	return n
}

func readArray(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		var x int
		fmt.Scan(&x)
		a[i] = x
	}
	return a
}

func printReversed(array []int) {
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Printf("%d ", array[i])
	}
}
