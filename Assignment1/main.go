package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MeirhanSyzdykov/Assignment1/Bank"
	"github.com/MeirhanSyzdykov/Assignment1/Company"
	"github.com/MeirhanSyzdykov/Assignment1/Library"
	"github.com/MeirhanSyzdykov/Assignment1/Shapes"
)

func library() {
	library := Library.NewLibrary()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Menu")
		fmt.Println("1. Add Book")
		fmt.Println("2. Borrow Book")
		fmt.Println("3. Return Book")
		fmt.Println("4. List Available Books")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter book title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter book author: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			id := library.NumForAutoIncrement
			library.NumForAutoIncrement += 1

			book := Library.Book{
				ID:         id,
				Title:      title,
				Author:     author,
				IsBorrowed: false,
			}

			library.AddBook(book)

		case "2":
			fmt.Print("Enter book ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			library.BorrowBook(id)

		case "3":
			fmt.Print("Enter book ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			library.ReturnBook(id)

		case "4":
			library.ListAvailableBooks()

		case "5":
			return

		default:
			fmt.Println("There is no such option. Please try again")
		}
	}
}

func shapes() {
	rectangle := Shapes.Rectangle{Length: 5, Width: 3}
	circle := Shapes.Circle{Radius: 4}
	square := Shapes.Square{Side: 2}
	triangle := Shapes.Triangle{SideA: 8, SideB: 7, SideC: 9}

	shapes := make([]Shapes.Shapes, 0)
	shapes = append(shapes, rectangle, circle, square, triangle)

	for _, shape := range shapes {
		fmt.Printf("%T: %v\n", shape, shape)
		fmt.Printf("Area: %.2f\n", shape.Area())
		fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
		fmt.Println()
	}
}

func company() {
	fullTimeEmployee1 := Company.FullTimeEmployee{
		ID:     1,
		Name:   "Arthur",
		Salary: 500,
	}

	partTimeEmployee1 := Company.PartTimeEmployee{
		ID:        2,
		Name:      "Prefect",
		HourlyPay: 50,
		Hours:     10,
	}

	company := Company.NewCompany()
	company.AddEmployee(fullTimeEmployee1.ID, fullTimeEmployee1)
	company.AddEmployee(partTimeEmployee1.ID, partTimeEmployee1)

	company.ListEmployees()
}

func bank() {
	for {
		bank := Bank.BankAccount{
			AccountNumber: "1234",
			AccountOwner:  "Scrooge",
			Balance:       0,
			Transactions:  make([]string, 0),
		}
		reader := bufio.NewReader(os.Stdin)

		for {
			fmt.Println("\nBank Menu")
			fmt.Println("1. Deposit")
			fmt.Println("2. Withdraw")
			fmt.Println("3. Get Balance")
			fmt.Println("4. Exit")
			fmt.Print("Choose an option: ")

			choice, _ := reader.ReadString('\n')
			choice = strings.TrimSpace(choice)

			switch choice {
			case "1":
				fmt.Print("Enter amount: ")
				input, _ := reader.ReadString('\n')
				amount, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
				bank.Deposit(amount)

			case "2":
				fmt.Print("Enter amount: ")
				input, _ := reader.ReadString('\n')
				amount, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
				bank.Withdraw(amount)

			case "3":
				bank.GetBalance()

			case "4":
				return

			default:
				fmt.Println("There is no such option. Please try again")
			}
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nMain Menu")
		fmt.Println("1. Library")
		fmt.Println("2. Shapes")
		fmt.Println("3. Company")
		fmt.Println("4. Bank")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			library()

		case "2":
			shapes()

		case "3":
			company()

		case "4":
			bank()

		case "5":
			return

		default:
			fmt.Println("There is no such option. Please try again")
		}
	}
}
