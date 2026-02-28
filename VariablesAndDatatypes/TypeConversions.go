package main;

import "fmt"

func main(){
	var i int = 45
	var f float64 = float64(i)
	var u uint = uint(f)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(u)
}