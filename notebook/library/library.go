//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	name string
	quantity int
	history LendHistory
}

type Member struct{
	name string
	books []Book
} 
 
type Library struct {
	books map[string]Book
	members map[string][]Member
}

type LendHistory struct {
	checkoutTime time.Time
	checkinTime time.Time
}

func checkout(library *Library, bookTitle string,borrower string){
	book, isFound := library.books[bookTitle]
	if isFound && library.books[bookTitle].quantity >= 1 {
		//udate book checkout time and quantity
		book.history.checkoutTime = time.Now()
		book.quantity -= 1

		//update library 
		library.books[bookTitle] = book
		library.members[borrower] = append(library.members[borrower], []Member{})
		fmt.Println(library)
		
		fmt.Println("Checkout",bookTitle,"successfully for",borrower)
		fmt.Println("Quantity available:",book.quantity)
	}else if !isFound{
		fmt.Println(bookTitle,"not belong to library")
	}else if book.quantity < 1{ 
		fmt.Println(bookTitle, "not available to checkout")
	}
}

func checkin(library *Library, bookTitle string){
	book, isFound := library.books[bookTitle]
	if isFound {
		book.history.checkinTime = time.Now()
		book.quantity += 1

		library.books[bookTitle] = book
		fmt.Println("Checkin",bookTitle,"successfully")
		
	}else{
		fmt.Println("Cannot checkin")
	}
	
}

func addBook(library *Library,bookTitle string,book Book){
	library.books[bookTitle] = book
}

func addMember(library *Library,member Member){
	library.members[member] = []string{}
}

func printLibInfo(library *Library){
	numOfBooks := len(library.books)
	numOfMembers := len(library.members)

	fmt.Println("There are",numOfBooks,"book titles and",numOfMembers,"members")

	fmt.Println("Books are:")
	for name,val := range library.books{
		fmt.Println(name,"with",val.quantity,"copies")
	}

	fmt.Println("Members are:")
	for member,books := range library.members{
		fmt.Println(member,"is borrowing",books)
	}
}

func main() {
	var (
		brave = Book{name:"Brave, not perfect",quantity: 2}
		ocean = Book{name:"On earth, we're briefly perfect",quantity: 39}
		habit = Book{name:"Atomic habit",quantity: 12}
		running = Book{name:"What I talk about when I talk about running",quantity: 8}
	)


	books := make(map[string]Book)
	members := make(map[Member][]string)
	library := Library{books,members}

	addBook(&library,"Brave, not perfect",brave)
	addBook(&library,"On earth, we're briefly perfect",ocean)
	addBook(&library,"Atomic habit",habit)
	addBook(&library,"What I talk about when I talk about running",running)

	addMember(&library,"Holly")
	addMember(&library,"Ben")
	addMember(&library,"Claire")

	checkout(&library,"Brave, not perfect",Holly)
	checkout(&library,"Brave, not perfect",Ben)
	checkout(&library,"On earth, we're briefly perfect",Claire)

	checkin(&library,"Brave, not perfect")
	checkin(&library,"Atomic habit")

	checkout(&library,"Brave, not perfect",Ben)
	checkout(&library,"Atomic habit",Ben)
	checkout(&library,"On earth, we're briefly perfect",Ben)
	checkout(&library,"Brave, not perfect",Claire)
	checkout(&library,"Nudge",Claire)

	printLibInfo(&library)
	//fmt.Println(library)
}

