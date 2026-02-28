package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumber() int {
	return rand.Intn(10)
}

func databaseConnectionAttempt() {
    for connectionAttempt := 3; connectionAttempt > 0; connectionAttempt-- {
        if randomNumber := generateRandomNumber(); randomNumber < 3 {
            fmt.Println("Connection Established")
            return 
        }
        fmt.Printf("Attempting retry %d \n", connectionAttempt)
    }
    fmt.Println("Alert: Database unreachable") 
}

func main() {
databaseConnectionAttempt()
}
