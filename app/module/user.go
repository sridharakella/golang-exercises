package module

import "fmt"

type User struct {
	Name string //access control
}

func sample() {
	var user User
	user.Name = "John Doe"
	fmt.Println(user.Name)

}
