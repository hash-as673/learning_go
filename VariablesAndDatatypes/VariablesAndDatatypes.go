package main;

import "fmt"

func main(){
	var greeting string = "Hello World"
	i := 34 
	f := 3.14
	b := true
	c := 3 + 4i
	fmt.Println(greeting)
	fmt.Println(i)
	fmt.Println(f)
	fmt.Println(b)
	fmt.Println(c)

	const(
		port = 8080
		host = "localhost"
	)

	fmt.Println(port)
	fmt.Println(host)
}