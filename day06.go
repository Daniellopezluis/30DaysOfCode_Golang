package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var T int
	fmt.Scanf("%d", &T)
	for T < 1 || T > 10 {
		fmt.Println("Please insert a valid number:")
		fmt.Scanf("%d", &T)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < T; i++ {
		_ = scanner.Scan()
		line := scanner.Text()
		var even, odd []rune
		nextEven := true
		for _, x := range line {
			if nextEven {
				even = append(even, x)
				nextEven = false
			} else {
				odd = append(odd, x)
				nextEven = true
			}
		}
		fmt.Println(string(even), string(odd))
	}
}
