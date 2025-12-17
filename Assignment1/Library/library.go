package Library

import "fmt"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[string]Book),
	}
}

func (l *Library) AddBook(id string, book Book) {
	for _, book := range l.Books {
		if book.ID == id {
			fmt.Println("Book with such ID is already in library")
			return
		}
	}

	l.Books[id] = book
	fmt.Println("Book added successfully")
}

func (l *Library) BorrowBook(id string) {
	book, ok := l.Books[id]

	if !ok {
		fmt.Println("Book not found")
		return
	}

	if book.IsBorrowed {
		fmt.Println("Book is already borrowed")
		return
	}

	book.IsBorrowed = true
	l.Books[id] = book
	fmt.Println("Book is successfuly borrowed")
}

func (l *Library) ReturnBook(id string) {
	book, ok := l.Books[id]

	if !ok {
		fmt.Println("Book not found")
		return
	}

	if !book.IsBorrowed {
		fmt.Println("Book is already returned")
		return
	}

	book.IsBorrowed = false
	l.Books[id] = book
	fmt.Println("Book is successfuly returned")
}

func (l *Library) ListAvailableBooks() {
	found := false

	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %v, Title: %v, Author: %v\n", book.ID, book.Title, book.Author)
			found = true
		}
	}

	if !found {
		fmt.Println("No available books")
	}
}
