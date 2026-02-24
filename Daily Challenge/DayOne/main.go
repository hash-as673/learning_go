package main

import (
	"fmt"
	"math/rand"
)

func rollStats()(strength int, agility int){
	strength = rand.Intn(20)
	agility = rand.Intn(20)
	return
}


func calculatePower(strength int, agility int) float64{
    power := (float64(strength) + float64(agility)) * 1.5
	return power
}

func main(){
	strength, agility := rollStats()
	power := calculatePower(strength, agility)
	fmt.Printf("Player Stats: %d STR, %d AGI. Total Combat Power: %.2f", strength, agility, power)
	
}