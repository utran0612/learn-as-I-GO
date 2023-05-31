//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these cirsumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import (
	"fmt"
	"math/rand"
)

func roleDice(sides, numOfDice int) int{
	var res int
	for i := 0;i<numOfDice;i++{
		res += rand.Intn(sides) + 1
	}
	return res
}

func diceRoll(times,numOfDice,sides int) {
	var sumOfDiceRolls int
	for i :=0 ;i < times;i++{
		newSum := roleDice(sides, numOfDice)
		fmt.Println("Roll #",i+1,":",newSum)
		sumOfDiceRolls += newSum
	}
	fmt.Println("Sum of the dice roll is", sumOfDiceRolls)
	if times == 2 && numOfDice == 7{
		fmt.Println("Snake eyes")
	}else if times == 7{
		fmt.Println("Lucky 7")
	}else if times % 2 == 0{
		fmt.Println("Even")
	}else if times % 2 == 1{
		fmt.Println("Odd")
	}
}

func main() {
	diceRoll(1,2,12)
}