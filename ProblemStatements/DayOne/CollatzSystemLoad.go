package main

import "fmt"

func collatz(n int) int {
	step := 0

	for n != 1 {
		if n%2 == 0 {
			n /= 2
			// fmt.Println("Step:",n)
		} else {
			n = n*3 + 1

		}
		step++
	}
	return step

}

func main() {
	const n = 10
	steps := collatz(n)
	fmt.Println("Steps:",steps)
}
