package library

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

func (b *Book) String() string {
	return fmt.Sprintf("ID: %d | Title: %s | Author: %s | Copies: %d", b.ID, b.Title, b.Author, b.Copies)
}
