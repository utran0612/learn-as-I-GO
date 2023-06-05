//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

type Part struct {
	content string
}

func printParts(slice []Part){
	for i:=0;i<len(slice);i++{
		part := slice[i]
		fmt.Println(part)
	}
}

func main() {
	x := Part{"this"}
	y := Part{"is"}
	z := Part{"a"}

	// create an assembly line having any 3 parts
	assemblyLine := []Part{x,y,z}
	fmt.Println("Newly created assembly line is:",assemblyLine)

	// add 2 new parts to the line 
	assemblyLine = append(assemblyLine, Part{"full"}, Part{"sentence"})
	fmt.Println("Added 2 new parts, here's the assembly line now:",assemblyLine)

	// Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[3:]
	fmt.Println("The assembly line with only 2 new parts",assemblyLine)

}