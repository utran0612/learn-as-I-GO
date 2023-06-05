//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Item struct {
	name string
	securityState bool
}

func activate(item *Item) {
	if !item.securityState {
		item.securityState = true
	}
}

func deactivate(item *Item){
	if item.securityState{
		item.securityState = false
	}
}

func checkout(slice []Item){
	// for _,item := range slice {
	// 	fmt.Println(item)
	// 	deactivate(&item)
	// }

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
		deactivate(&slice[i])
	}
}

func main() {

	// Create at least 4 items, all with active security tags
	iphone := Item{"iphone 4",true}
	gate := Item{"auto",true}
	speaker := Item{"sony",true}
	port := Item{"80",true}

	// Store them in a slice or array
	slice := []Item{iphone,gate,speaker,port}

	//  - Deactivate any one security tag in the array/slice
	deactivate(&slice[0])

	//  - Call the checkout() function to deactivate all tags
	checkout(slice)

	//  - Print out the array/slice after each change
	fmt.Println(slice)
}


