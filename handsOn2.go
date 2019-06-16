package main

import "fmt"

type person struct {

	fname string
	lname string
	age int

}

type secretAgent struct {

	person
	kill bool

}

func (p	person) pSpeak() {

	fmt.Println("Hi people! I'm a person =)")

}

func (s	secretAgent) saSpeak() {

	fmt.Println("Hi dude! A secret agente here!")

}

func main () {

	p1 := person {
		"John",
		"Whatson",
		32,
	}

	sa1 := secretAgent {
		person{
			"Sherlok",
			"Holmes",
			36,
		},
		true,
	}

	fmt.Println(p1.fname)
	fmt.Println(sa1.fname)

	p1.pSpeak()
	sa1.saSpeak()
	
}