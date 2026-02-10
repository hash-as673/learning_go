package main

import (
	"fmt"

)

var i,j int = 1,2 
func main(){
	// var i,j int = 1,2 not allowed here
	var c,python,java = false, false, true
	k := 3 // Short variable declaration, only inside functions
	fmt.Println(i,j,c,python,java)
	fmt.Println(k)

	//Short variable declaration


	//Datatypes
	var i int = 33
	var f float32 =33.3
	var f2 float64 = 3423.234
	var by byte = 12
	var str string = "Hello World"
	var cm complex64 = 1 + 2i
	fmt.Println(i,f,f2,by,str,cm,len(str))

	const PI = 3.14
	fmt.Println(PI)


// 	The zero value is:

// 0 for numeric types,
// false for the boolean type, and
// "" (the empty string) for strings.
}