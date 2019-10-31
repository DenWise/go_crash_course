package main

import "fmt"

func main() {
	// Array
	var fruitArray [2]string
	
	//Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	// Declare and assign
	thingArray := [2]string{"Something","Anything"}

	fmt.Println(fruitArray)
	fmt.Println(thingArray)

	// Slice
	fruitSlice := []string{"Apple","Orange","Pineapple"}

	rangeBottom, rangeTop := 1, 3

	fmt.Println(fruitSlice,len(fruitSlice),fruitSlice[rangeBottom:rangeTop])
}