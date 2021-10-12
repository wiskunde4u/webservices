package main

import (
	"fmt"

	"github.com/wiskunde4u/webservice/models"
)

func main() {
	//	fmt.Println("Hello all")

	u := models.User{ //same working with models package as with the fmt package.
		ID:        2,
		FirstName: "Tricia",
		LastName:  "McMillan", //multi line initialization syntax for this user type "comma" for the final line.
	}
	fmt.Println(u) // format will space out some things making it easier to read
}

// create another package in the module space, being descrete unit of codes thta belongs to concepts
// so we create a package that is a folder called  MODELS
// the name is inported ... so we have to update the go.mod
// directory to be used is "models"

//to execut you can use the following command
//go run github.com/wiskunde4u/webservice/
