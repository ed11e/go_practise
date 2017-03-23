package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone int
}

type Student struct {
	Human
	university string
	course     string
}

type Grad struct {
	Human
	job string
}

func (h *Human) SayHi() {
	fmt.Println("Hi! My name is %s\n", h.name)
}

func (h *Human) GiveAge() {
	fmt.Println("I am %s\n", h.age)
}

func (h *Human) GiveNumber() {
	fmt.Println("My phone number is %s\n", h.phone)
}

func (g *Grad) SayHi() {
	fmt.Println("Hi! My name is %s\n", g.name, "I am a grad")
}
