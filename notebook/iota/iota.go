//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

type Operator int

const(
	Add Operator = iota
	Subtract
	Multiply
	Divide
)

func (o Operator) operations(a,b int) int {
	switch o {
	case Add:
		return a+b
	case Subtract:
		return a-b
	case Multiply:
		return a*b
	case Divide:
		return a/b
	default:
		return 0
	}
}

func main() {
	fmt.Println(Add.operations(3,4))
	fmt.Println(Subtract.operations(63,12))
	fmt.Println(Multiply.operations(2,14))
	fmt.Println(Divide.operations(930,3))
}

