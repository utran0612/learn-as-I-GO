//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import (
	"fmt"
	"math"
)

type Coordinate struct {
	x, y float64
}


type Rectangle struct {
	topLeft, bottomRight Coordinate
}

func calculateRectangleArea(rectangle Rectangle) {
	width := math.Abs(rectangle.topLeft.x - rectangle.bottomRight.x)
	length := math.Abs(rectangle.topLeft.y - rectangle.bottomRight.y)

	fmt.Println("Width of rectangle is:",width)
	fmt.Println("Length of rectangle is:",length)

	area := math.Abs(width*length)
	fmt.Println("The area of the rectangle is:",area)
}


func calculateRectanglePerimeter(rectangle Rectangle){
	width := math.Abs(rectangle.topLeft.x - rectangle.bottomRight.x)
	length := math.Abs(rectangle.topLeft.y - rectangle.bottomRight.y)

	peri := (width + length)*2
	fmt.Println("The perimeter of the rectangle is:",peri)

}

func main(){
	topLeft := Coordinate{0,7}
	bottomRight := Coordinate{10,0}
	rec := Rectangle{topLeft,bottomRight}
	calculateRectangleArea(rec)
	calculateRectanglePerimeter(rec)
	rec.topLeft.x *= 2
	rec.bottomRight.x *= 2
	rec.topLeft.y *= 2
	rec.bottomRight.y *= 2
	calculateRectangleArea(rec)
	calculateRectanglePerimeter(rec)
}




