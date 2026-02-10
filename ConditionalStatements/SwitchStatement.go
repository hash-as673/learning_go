package main

import "fmt"

func main(){
	rating := 4

	switch rating {
	case 1:
		fmt.Println("Bad")
	case 2:
		fmt.Println("Not Good")
	case 3:
		fmt.Println("Average")
	case 4:
		fmt.Println("Excellent")
	case 5:
		fmt.Println("Outstanding")
	default:
		fmt.Println("Invalid Input")
	}
}