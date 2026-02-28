package main;

import "fmt"

func main(){
	i:=4
	switch(i){
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Some other number")
	} 
}