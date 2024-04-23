package main

import "fmt"

type ContactInfo struct {
	email string
	zip   int32
}

type Person struct {
	firstname string
	lastName  string
	ContactInfo
}

func main() {
	jim := Person{
		firstname: "Jim",
		lastName:  "Morrison",
		ContactInfo: ContactInfo{
			email: "morri@doors.com",
			zip:   111111,
		},
	}

	jim.updateName("Oliver")
	jim.printName()
}

func (p Person) printName() {
	fmt.Printf("%s %s", p.firstname, p.lastName)
}

func (p *Person) updateName(newFirstName string) {
	(*p).firstname = newFirstName
}
