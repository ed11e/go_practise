package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
	Age                 int
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetAge() int {
	return u.Age
}

func main() {
	u := &User{"John", "Smith", 30}
	fmt.Println(u.GetFirstName())
	fmt.Println(u.GetLastName())
	fmt.Println(u.GetAge())
}
