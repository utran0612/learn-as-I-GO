//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//

package main

import (
	"fmt"
	"math/rand"
)

//--Requirements:
//* Write a function that accepts a person's name as a
//  parameter and displays a greeting to that person.
func greetingName (name string) {
	fmt.Println("Hello", name,"!")
}
//* Write a function that returns any message, and call it from within
//  fmt.Println()
func returnMessage (message string) string {
	return message
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func addThreeNumbers (a,b,c int) int {
	return a+b+c
}
//* Write a function that returns any number
func returnNumber() int {
	return rand.Int()
}
//* Write a function that returns any two numbers
func returnAnyTwoNumbers(a, b int) (int, int){
	return a,b
}
//* Add three numbers together using any combination of the existing functions.
//  - Print the result
func addThree() {
	a, b, c := returnNumber(), returnNumber(), returnNumber()
	fmt.Println("3 numbers are",a,b,c)
	res := addThreeNumbers(a,b,c)
	fmt.Println(res)
}
//* Call every function at least once
func main() {
	greetingName("Uyen")
	fmt.Println(returnMessage("breeze"))
	returnNumber()
	fmt.Println(returnAnyTwoNumbers(returnNumber(),returnNumber()))
	addThree()
}
