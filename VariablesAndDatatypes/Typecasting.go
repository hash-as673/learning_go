package main

import "fmt"

func main(){
	i:= 45
	f:= float64(i)
	var u uint = uint(f)
	fmt.Println(i,f,u)
}