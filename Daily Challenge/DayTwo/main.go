package main

import (
	"fmt"
	"math/rand"
)

func initiateBossFight(bossHP int){

	for bossHP > 0{

		switch randomAttackPower(){
		case 0:
			fmt.Println("Miss!")
		case 1:
			fmt.Println("Light Attack!")
			bossHP -= 5
			fmt.Println("Boss HP:",bossHP)
		case 2:
			fmt.Println("Heavy Attack!")
			bossHP -= 15
			fmt.Println("Boss HP:",bossHP)
		case 3:
			fmt.Println("Critical Hit!")
			bossHP -= 25
			fmt.Println("Boss HP:",bossHP)
		default:
			fmt.Println("Error: Invalid attack power!")
		}

	}
	fmt.Println("Boss is defeated!") 
}

func randomAttackPower()(a int){
	a = rand.Intn(4)
	return
}


func main(){
	defer fmt.Println("System Log: Boss encounter concluded.")

	fmt.Println("Initiating Boss Fight!")
	initiateBossFight(100)

}