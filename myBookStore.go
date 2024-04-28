package main

import (
	"fmt"
	"time"
)

type Book struct {
	title    string
	author   string
	numPages int
	price    float64
}


type Customer struct {
	name          string
	selectedBooks []Book
	selectedBooksQty []int
	totalCost float64
}

func clearScreen() {
	fmt.Print("\033[2J")
}

// func createBook() (newBook Book) {

// 	fmt.Print("Name of book: ")
// 	fmt.Scanln(&newBook.title)

// 	fmt.Print("Author of book: ")
// 	fmt.Scanln(&newBook.author)

// 	fmt.Print("Total number of pages: ")
// 	fmt.Scanln(&newBook.numPages)

// 	fmt.Print("Price of book: ")
// 	fmt.Scanln(&newBook.price)

// 	return newBook
// }

func printMenu() (choice int) {

	fmt.Print("\n\t(1) add new books")
	fmt.Print("\n\t(2) buy books")
	fmt.Print("\n\t(3) exit: ")
	fmt.Scan(&choice)

	return choice

}

func showBooks(books []Book) {
	for i := 0; i < len(books); i++ {
		fmt.Printf("\n(%v)%v\t\t$%.2f", i, books[i].title, books[i].price)
	}
}

func addBooks(myBooks *[]Book) {
	var newBook Book
	var anotherBook string

	for {

		fmt.Print("\nBook Title: ")
		fmt.Scanln(&newBook.title)

		fmt.Print("\nBook Author: ")
		fmt.Scanln(&newBook.author)

		fmt.Print("\nNumber of pages in Book: ")
		fmt.Scanln(&newBook.numPages)

		fmt.Print("\nBook Price: ")
		fmt.Scanln(&newBook.price)

		*myBooks = append(*myBooks, newBook)

		fmt.Print("\nAdd another books? (y/n): ")
		fmt.Scanln(&anotherBook)

		if anotherBook == "n" {
			break
		}

	}
}

func buyBook(myBooks []Book) {

	var customers []Customer
	var bookOfChoice int
	var qty int
	var anotherCustomer string
	var anotherBook string
	var customer Customer

	// for each customers
	for {

		clearScreen()

		fmt.Print("\nEnter customer name: ")
		fmt.Scan(&customer.name)

		// for each customer, they can buy how many times they want (for each selectedBooks)
		for {

			clearScreen()
			showBooks(myBooks)

			fmt.Print("\nPlease select a book to buy: ")
			fmt.Scan(&bookOfChoice)

			fmt.Print("\nHow many you wanna buy?: ")
			fmt.Scan(&qty)

			if bookOfChoice > len(myBooks) || bookOfChoice < 0 {
				fmt.Print("\nPlease choose another book.")
			} else {
				customer.selectedBooks = append(customer.selectedBooks, myBooks[bookOfChoice])
				customer.selectedBooksQty = append(customer.selectedBooksQty, qty)
			}

			fmt.Print("\nChoose another book to buy? (y/n): ")
			fmt.Scan(&anotherBook)

			if anotherBook == "n" {
				break
			}

		}

		for i:=0; i<len(customer.selectedBooks); i++ {
			customer.totalCost += customer.selectedBooks[i].price * float64(customer.selectedBooksQty[i])
		}

		// show total cost
		clearScreen()
		fmt.Printf("\nCustomer Name: %s\t\t\tDate: %s", customer.name, time.Now().Format("02-Jan-2006"))
		fmt.Print("\n-----------------------------------------------------")
		showBooks(customer.selectedBooks)
		fmt.Print("\n-----------------------------------------------------")
		fmt.Printf("\nTotal Cost: \t\t\t$%.3f",customer.totalCost)

		customers = append(customers, customer)

		fmt.Print("\nAnother Customer? (y/n): ")
		fmt.Scan(&anotherCustomer)

		if anotherCustomer == "n" {
			break
		}

	}

	// show the customer list and their total selcted books
	for i := 0; i < len(customers); i++ {
		fmt.Printf("\n(%v)\t%s\t%v", i, customers[i].name, len(customers[i].selectedBooks))
	}

}

func main() {

	var myBooks []Book
	var demoBooks = []Book {
		{
			title: "Harry Potter(1)",
			author: "JK.Rowlling",
			numPages: 200,
			price: 150.55,
		},
		{
			title: "Harry Potter(2)",
			author: "JK.Rowlling",
			numPages: 130,
			price: 130.55,
		},
		{
			title: "Harry Potter(3)",
			author: "JK.Rowlling",
			numPages: 250,
			price: 250.55,
		},
	}
	var another string

	clearScreen()

	fmt.Println("\n\n\t\tWelcome from book store")
	fmt.Println("\t\t-----------------------")

	for {

		choice := printMenu()

		if choice == 1 {
			addBooks(&myBooks)
		}

		if choice == 2 {
			buyBook(demoBooks)
		}

		fmt.Print("\nChoose menu again? (y/n): ")
		fmt.Scan(&another)
		if another == "n" {
			break
		}
		clearScreen()
	}
}
