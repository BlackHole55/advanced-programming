package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MeirhanSyzdykov/Assignment1/Agents"
	"github.com/MeirhanSyzdykov/Assignment1/Bank"
	"github.com/MeirhanSyzdykov/Assignment1/Company"
	"github.com/MeirhanSyzdykov/Assignment1/Library"
	"github.com/MeirhanSyzdykov/Assignment1/Shapes"
	"github.com/MeirhanSyzdykov/Assignment1/Tickets"
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

	company := Company.NewCompany()

	fullTimeEmployee1 := Company.FullTimeEmployee{
		ID:     company.NumForAutoIncrement,
		Name:   "Arthur",
		Salary: 500,
	}
	company.NumForAutoIncrement += 1

	partTimeEmployee1 := Company.PartTimeEmployee{
		ID:        company.NumForAutoIncrement,
		Name:      "Prefect",
		HourlyPay: 50,
		Hours:     10,
	}
	company.NumForAutoIncrement += 1

	company.AddEmployee(fullTimeEmployee1)
	company.AddEmployee(partTimeEmployee1)

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

func store() {
	store := Tickets.NewTicketStore()
	agents := make(map[string]Agents.Agent)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Menu")
		fmt.Println("1. Create Ticket")
		fmt.Println("2. Add Agent")
		fmt.Println("3. Assign Ticket to Agent")
		fmt.Println("4. Resolve Ticket")
		fmt.Println("5. List All Tickets")
		fmt.Println("6. List OPEN Tickets")
		fmt.Println("7. List DONE Tickets")
		fmt.Println("8. List Unassigned Tickets")
		fmt.Println("9. Exit")

		fmt.Print("Choose an option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter ticket ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Enter ticket Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter ticket Description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			priorityVal := 0

			for {
				fmt.Print("Enter ticket Priority: ")
				priority, _ := reader.ReadString('\n')
				priority = strings.TrimSpace(priority)
				priorityVal, _ = strconv.Atoi(priority)

				if priorityVal == 0 {
					fmt.Println("Priority must be between 1 and 3")
					continue
				}

				break
			}

			ticket := Tickets.Ticket{
				ID:          id,
				Title:       title,
				Description: description,
				Priority:    priorityVal,
				Status:      "OPEN",
			}

			store.Create(ticket)

		case "2":
			fmt.Print("Enter Agent Type (Human/Bot): ")
			agentType, _ := reader.ReadString('\n')
			agentType = strings.TrimSpace(agentType)

			if agentType == "Human" {
				fmt.Print("Enter Agent ID: ")
				id, _ := reader.ReadString('\n')
				id = strings.TrimSpace(id)

				fmt.Print("Enter Agent Name: ")
				name, _ := reader.ReadString('\n')
				name = strings.TrimSpace(name)

				agent := Agents.HumanAgent{
					ID:   id,
					Name: name,
				}

				agents[id] = agent

			} else if agentType == "Bot" {
				fmt.Print("Enter Agent ID: ")
				id, _ := reader.ReadString('\n')
				id = strings.TrimSpace(id)

				fmt.Print("Enter Agent Name: ")
				name, _ := reader.ReadString('\n')
				name = strings.TrimSpace(name)

				fmt.Print("Enter Agent Version: ")
				version, _ := reader.ReadString('\n')
				version = strings.TrimSpace(version)

				agent := Agents.BotAgent{
					ID:      id,
					Name:    name,
					Version: version,
				}

				agents[id] = agent
			} else {
				fmt.Println("Input must be either Human or Bot")
			}

		case "3":
			fmt.Print("Enter Ticket ID: ")
			ticketId, _ := reader.ReadString('\n')
			ticketId = strings.TrimSpace(ticketId)

			ticketExists := false

			for key := range store.Items {
				if key == ticketId {
					ticketExists = true
				}
			}

			fmt.Print("Enter Agent ID: ")
			agentId, _ := reader.ReadString('\n')
			agentId = strings.TrimSpace(agentId)

			agentExists := false

			for key := range agents {
				if key == agentId {
					agentExists = true
				}
			}

			if !ticketExists {
				fmt.Println("Ticket does not exists")
				break
			} else if !agentExists {
				fmt.Println("Agent does not exist")
				break
			}

			store.Assign(ticketId, agentId)

		case "4":
			fmt.Print("Enter Ticket ID: ")
			ticketId, _ := reader.ReadString('\n')
			ticketId = strings.TrimSpace(ticketId)

			store.Resolve(ticketId)

		case "5":
			store.ListAll()

		case "6":
			store.ListByStatus("OPEN")

		case "7":
			store.ListByStatus("DONE")

		case "8":
			store.ListUnassigned()

		case "9":
			return

		default:
			fmt.Println("There is no such option. Please try again")
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
		fmt.Println("5. Ticket Store")
		fmt.Println("6. Exit")
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
			store()

		case "6":
			return

		default:
			fmt.Println("There is no such option. Please try again")
		}
	}
}
