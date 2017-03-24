package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
	Age                 int
	Employed            bool
	BankBalance         float64
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

func (u *User) Birthday() int {
	u.Age++
	return u.Age
}

func (u *User) Retire() bool {
	u.Employed = false
	return u.Employed
}

func (u *User) NewJob() bool {
	u.Employed = true
	return u.Employed
}

func (u *User) GetPaid(i float64) float64 {
	u.BankBalance += i
	return u.BankBalance
}

func (u *User) SpendMoney(i float64) float64 {
	u.BankBalance -= i
	return u.BankBalance
}

func (u *User) GetBankBalance() float64 {
	return u.BankBalance
}

func main() {
	u := &User{"John", "Smith", 30, false, 10}
	fmt.Println(u.GetFirstName())
	fmt.Println(u.GetLastName())
	fmt.Println(u.GetAge())
	fmt.Println(u.GetEmployed())
	fmt.Println("User getting older")
	u.Birthday()
	fmt.Println(u.GetAge())
	fmt.Println("Got a Job!!!")
	u.NewJob()
	fmt.Println(u.GetEmployed())
	fmt.Println("Time to Retire!!!")
	u.Retire()
	fmt.Println(u.GetEmployed())
	u.GetPaid(20.50)
	fmt.Println(u.GetBankBalance())

}
