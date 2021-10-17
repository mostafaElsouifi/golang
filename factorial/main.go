package main

import "fmt"

func factorial(n int) int {
	sum := 1
	for i := 1; i <= n; i++ {
		sum = sum * i
	}
	return sum
}
func main() {
	fmt.Println(factorial(5))
}
