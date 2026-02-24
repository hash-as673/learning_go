package main;

import "fmt"

func main(){
	var i int = 45
	var f float64 = float64(i)
	var u uint = uint(f)
	var c,python,java = true, false, true
	var b byte = 10
	var cn complex64 = 10+5i;

	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(cn)
	fmt.Println(u)
	fmt.Println(c,python,java)


	// Variable Short form
	j := 10
	s := "Hello World"
	fmt.Println(j,s)

	const (
		port=8080
		host="localhost"
	)


}