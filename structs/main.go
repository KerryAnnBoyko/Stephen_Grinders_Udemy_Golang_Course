package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	fmt.Println("Initial")
	jim.print()

	jim.updateNameIncorrectly("jimmy-no-pointer")
	fmt.Println("didn't work because it's not a pointer")
	jim.print()

	// &jim is the pointer slash memory address. The type is not int, (as in C) but rather *person! SO much better than C!
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")

	(&jim).updateName("Jimmy")

	fmt.Println("firstName should now be 'Jimmy'")
	jim.print()

	// Here's a weird thing. Okay. So, we have defined a reciever called "updateName" which works on type *person
	// However, we can now call that reciever on type person (no asterisk!) and it will automatically assume you
	// want the pointer. Nice little shortcut!

	jim.updateName("Jimbo")

	fmt.Println("firstName should now be 'Jimbo'")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// this won't work. We're passing p as a value, not as a pointer to the struct.
func (p person) updateNameIncorrectly(newFirstName string) {
	p.firstName = newFirstName
	fmt.Println("This is a local reference, this doesn't change it")
	p.print()
}

// The asterisk tells us that we're getting a *pointer* to an existing person, NOT a copy of a new person.
func (p *person) updateName(newFirstName string) {
	// if we try to log the pointer, we don't get the random hex address, we get "&" plus the value of the element!
	// GO ROCKS!
	fmt.Printf("Hello! I am a pointer! My name is %v\n", p)
	// *p means "the element referenced at the pointer 'p'"
	(*p).firstName = newFirstName
}
