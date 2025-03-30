package main

import (
	"fmt"
	"library-manage-app/library"
	"library-manage-app/utils"
	"sync"
	"time"
)

func main() {
	lib := library.NewLibrary()

	// Adding Sample Books
	lib.AddBook(1, "The Go Programming Language", "Alan Donovan", 3)
	lib.AddBook(2, "Concurrency in Go", "Katherine Cox-Buday", 2)

	// Adding Sample Users
	lib.AddUser(101, "Alice")
	lib.AddUser(102, "Bob")

	// Display initial state
	lib.ListBooks()
	lib.ListUsers()

	// Simulating concurrent borrowing
	var wg sync.WaitGroup

	borrowBook := func(userID, bookID int) {
		defer wg.Done()
		lib.BorrowBook(userID, bookID)
		utils.Log(fmt.Sprintf("User %d borrowed book %d", userID, bookID))
	}

	returnBook := func(userID, bookID int) {
		defer wg.Done()
		time.Sleep((5 * time.Second))
		lib.ReturnBook(userID, bookID)
		utils.Log(fmt.Sprintf("User %d returned book %d", userID, bookID))
	}

	wg.Add(4)

	go borrowBook(101, 1)
	go borrowBook(102, 1)
	go borrowBook(101, 2)
	go returnBook(101, 1)

	wg.Wait()

	// Display final state
	lib.ListBooks()
}
