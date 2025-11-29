package main

import "fmt"

func main() {
	age := 10

	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	if age := 20; age >= 18 {
	fmt.Println("You are an adult")
} else {
	fmt.Println("You are a minor")
}
}

