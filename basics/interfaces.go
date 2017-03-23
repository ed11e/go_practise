package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
	Age                 int
	Employed            bool
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

func (u *User) GetEmployed() bool {
	return u.Employed
}

func main() {
	u := &User{"John", "Smith", 30, true}
	fmt.Println(u.GetFirstName())
	fmt.Println(u.GetLastName())
	fmt.Println(u.GetAge())
	fmt.Println(u.GetEmployed())
}
