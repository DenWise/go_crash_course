package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	//Assign kv
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Delete"] = "to_delete@gmail.com"


	fmt.Println(emails,len(emails))

	delete(emails,"Delete")
	fmt.Println(emails)

	emailsDefinedWithAssignment := map[string]string{"Bob":"bob@gmail.com","Sharon":"sharoon@gmail.com"}

	fmt.Println(emailsDefinedWithAssignment)

	namesDefinedWithAssignment := map[int]string{1:"Jake",2:"Jones",5:"Bill"}

	fmt.Println(namesDefinedWithAssignment)
	delete(namesDefinedWithAssignment,2)
	fmt.Println(namesDefinedWithAssignment,namesDefinedWithAssignment[5])
}