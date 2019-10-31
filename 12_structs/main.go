package main

import (
	"fmt"
	"strconv"
)

const male, female = "m","f"

// Define struct
type Person struct {
	firstName, lastName, city, gender string
	age int
}

// Define methods (value reciever)
func (p Person) greet() string {
	return "Hello, I'm " + p.firstName + " and I'm " + strconv.Itoa(p.age)
}

// Define methods (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) getMarried(newLastName string) {
	if p.gender == male {
		return
	} else {
		p.lastName = newLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Jenn", lastName: "Johnson", city:"San Francisco", gender: "f", age: 35}

	// Alternative
	person2 := Person{"Bob","Ford","Boston","m",25}

	persons := map[int]Person{1:person1,2:person2}

	for _, p := range persons {
		p.hasBirthday()
		p.getMarried("Riddle")
		fmt.Printf("%s %s is living in %s.\n",p.firstName,p.lastName,p.city)
		fmt.Println(p.greet())
	}
}