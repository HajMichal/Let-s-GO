package main

import "fmt"

type Contact struct {
	email 	string
	zipcode string
}
type Person struct {
	name 	string
	surname string
	contact Contact
}

func main() {
	var person = Person{
		name: "Michał",
		surname: "Haj",
		contact: Contact{
			email: "michalhaj@gmail.com",
			zipcode: "01-864",
		},
	}
	var person2 = Person{
		name: "Mariusz",
		surname: "Telok",
		contact: Contact{
			email: "MariuszekTeloczek@gmail.com",
			zipcode: "01-864",
		},
	}

	person.updateName("Klaudia")
	person.print()
	person2.print()

}


func (p *Person) updateName(newFirstName string) {
	p.name = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}