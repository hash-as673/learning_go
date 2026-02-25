package main;

import "fmt"

func main(){
	rating := 4
	
	switch rating {
	case 1:
		fmt.Println("Bad")
	case 2:
		fmt.Println("Average")
	case 3: 
		fmt.Println("Great")
	case 4:
		fmt.Println("Excellent")
	case 5:
		fmt.Println("Outstanding")
	default:
		fmt.Println("Invalid Rating")
	}
}