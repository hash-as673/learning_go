package main

import "fmt"

func main() {
	var rating = 4
	switch rating {
	case 1:
		fmt.Println("Bad")
	case 2:
		fmt.Println("Average")
	case 3:
		fmt.Println("Good")
	case 4:
		fmt.Println("Very Good")
	case 5:
		fmt.Println("Excellent")
	}

	//Type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("I am an integer and my value is %d\n", t)
		case string:
			fmt.Printf("I am a string and my value is %s\n", t)
		default:
			fmt.Printf("Unknown type\n")
		}
	}

	whoAmI(42)
	whoAmI("Golang")
	whoAmI(3.14)
}
