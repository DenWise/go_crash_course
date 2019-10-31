package main

import "fmt"

func main() {
	a := 5 
	b := &a
	fmt.Println(a,b) // 5(a) 0xc000012078(b)
	fmt.Printf("Type of b is %T\n",b) // *int - pointer to integer value
	fmt.Println(*b) // read val from address
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}