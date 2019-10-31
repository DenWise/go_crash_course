package main

import "fmt"

func main() {
	// Define slice
	ids := []int{33,76,59,48,19,23}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n",i,id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together

	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum", sum)

	// Range with map
	emailsDefinedWithAssignment := map[string]string{"Bob":"bob@gmail.com","Sharon":"sharoon@gmail.com"}

	for k, v := range emailsDefinedWithAssignment {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emailsDefinedWithAssignment {
		fmt.Println("Name: " + k)
	}

	for _,v := range emailsDefinedWithAssignment {
		fmt.Println("Email: " + v)
	}
}