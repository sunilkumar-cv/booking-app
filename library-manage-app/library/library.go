package library

import (
	"fmt"
	"sync"
)

type Library struct {
	Books map[int]*Book
	Users map[int]*User
	mu    sync.Mutex
}

func NewLibrary() *Library {
	return &Library{
		Books: make(map[int]*Book),
		Users: make(map[int]*User),
	}
}

// Add a new book to the library
func (lib *Library) AddBook(id int, title, author string, copies int) {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	lib.Books[id] = &Book{ID: id, Title: title, Author: author, Copies: copies}
	fmt.Println("Book added:", lib.Books[id])
}

// Add a new user to the library
func (lib *Library) AddUser(id int, name string) {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	lib.Users[id] = &User{ID: id, Name: name}
	fmt.Println("User added:", lib.Users[id])
}

// Borrow a book from the library
func (lib *Library) BorrowBook(userID, bookID int) {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	user, userExists := lib.Users[userID]
	book, bookExists := lib.Books[bookID]

	if !userExists {
		fmt.Println("User does not exist")
		return
	}
	if !bookExists {
		fmt.Println("Book does not exist")
		return
	}
	if book.Copies <= 0 {
		fmt.Println("No copies available for:", book.Title)
	}
	book.Copies--
	fmt.Printf("%s borrowed %s\n", user.Name, book.Title)
}

// Return a book to the library
func (lib *Library) ReturnBook(userID, bookID int) {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	book, bookExists := lib.Books[bookID]

	if !bookExists {
		fmt.Println("Book does not exist in the library")
		return
	}

	book.Copies++
	fmt.Printf("Book ID %d returned. Available copies: %d\n", book.ID, book.Copies)
}

// Display all books in the library
func (lib *Library) ListBooks() {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	fmt.Println("\nLibrary Books:")
	for _, book := range lib.Books {
		fmt.Println(book)
	}
}

// Display all users in the library
func (lib *Library) ListUsers() {
	lib.mu.Lock()
	defer lib.mu.Unlock()

	fmt.Println("\nLibrary Users:")
	for _, user := range lib.Users {
		fmt.Println(user)
	}
}
