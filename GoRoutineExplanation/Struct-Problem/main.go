package main

import "fmt"

// problem in go 1.22
type User struct {
	Name  string
	Email string
	//	Metadata [8192]byte
}

func main() {

	//in order to fix it
	// option 1
	/*

		users := []*{
			{Name: "sridhar", Email: "test@test.com"},
			{Name: "venkat", Email: "test@test.com"},
			{Name: "venkat1", Email: "1test@test.com"},
		}
	*/
	// option 2
	/*
		for i:= range users {

				fmt.Println("%p\n", &users[i])

			}
	*/

	users := []User{
		{Name: "sridhar", Email: "test@test.com"},
		{Name: "venkat", Email: "test@test.com"},
		{Name: "venkat1", Email: "1test@test.com"},
	}
	for i, user := range users {
		fmt.Printf("%p\n", &user)
		fmt.Println("%p\n", &users[i])

	}
}
