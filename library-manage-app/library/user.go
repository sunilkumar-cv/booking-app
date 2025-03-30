package library

import "fmt"

type User struct {
	ID   int
	Name string
}

func (u *User) String() string {
	return fmt.Sprintf("User ID: %d | Name: %s", u.ID, u.Name)
}
