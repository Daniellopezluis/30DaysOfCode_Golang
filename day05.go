package main

import "fmt"

func main() {
	var n, i int
	fmt.Scanf("%d", &n)
	for n < 2 || n > 20 {
		fmt.Println("Invalid value, please try again:")
		fmt.Scanf("%d", &n)
	}
	for i = 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}
