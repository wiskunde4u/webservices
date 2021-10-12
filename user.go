package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // Define a slice that is a collection of users, holding pointers to user objects.  We can manipulate the user withoug copying the data
	nextID = 1     // in the struct the compiler will be interpreted as an integer.

)
