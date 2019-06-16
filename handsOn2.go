// HANDS ON 2
// create a struct that holds person fields
// create a struct that holds secret agent fields and embeds person type
// attach a method to person: pSpeak
// attach a method to secret agent: saSpeak
// create a variable of type person
// create a variable of type secret agent
// print a field from person
// run pSpeak attached to the variable of type person
// print a field from secret agent
// run saSpeak attached to the variable of type secret agent
// run pSpeak attached to the variable of type secret agent

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