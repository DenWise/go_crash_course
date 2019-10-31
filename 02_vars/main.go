package main

import "fmt"

var age int32 = 25

func main() {
	// Main types
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 - unsigned
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	//var name = "Varien"
	//var age int32 = 25
	const isCool = true

	// Shorthand
	//name := "Varien"

	name, email := "Varien", "varien@gmail.com"

	fmt.Println(name, age, isCool, email)
	fmt.Printf("type of name is %T\n",name)
	fmt.Printf("type of age is %T\n",age)
	fmt.Printf("type of isCool is %T\n",isCool)
}