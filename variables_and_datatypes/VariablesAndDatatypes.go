package main

import "fmt"

func main() {
	// var name string = "Golang"
	i := 42
	f := 3.14
	s := "Hello Go!"
	b := true

	const pi = 3.14159

	fmt.Println("Integer:", i)
	fmt.Println("Float:", f)
	fmt.Println("String:", s)
	fmt.Println("Boolean:", b)
	fmt.Println("Constant Pi:", pi)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println("Server Host:", host)
	fmt.Println("Server Port:", port)

}
