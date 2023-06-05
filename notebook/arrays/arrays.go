//-Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	price int
	name string
}

func main() {
	cup := Product{5,"Cup"}
	speaker := Product{45,"bluetooth"}
	laptop := Product{2000,"macbook"}
	suitcase := Product{300,"samsonite"}

	shoppingList := [4]Product{cup,speaker,laptop}

	//last item on the list
	fmt.Println("Last item on the list is:",shoppingList[2].name)

	//total number of items
	fmt.Println("Total number of items:",len(shoppingList))

	//total cost of the items
	total := 0
	for i := 0; i<len(shoppingList);i++{
		item := shoppingList[i]
		total += item.price
	}
	fmt.Println("Total cost of the items:",total)

	//add the fourth item
	shoppingList[3] = suitcase
	//last item on the list
	fmt.Println("Last item on the list is:",shoppingList[len(shoppingList)-1].name)

	//total number of items
	fmt.Println("Total number of items:",len(shoppingList))

	//total cost of the items
	total = 0
	for i := 0; i<len(shoppingList);i++{
		item := shoppingList[i]
		total += item.price
	}
	fmt.Println("Total cost of the items:",total)

	
}