package main

import "fmt"


func main(){
	// Arrays are of fixed lengths
	// var fruitArray [2]string
	// // Assign Values
	// fruitArray[0] = "Apple"
	// fruitArray[1] = "Orange"

	// Declare and assign 
	// fruitArray := [2]string {"Apple", "Orange"}

	// fmt.Println(fruitArray);

	//Slices are of variable lengths
	fruitSlice := []string {"Apple", "Orange", "Grape"}
	fmt.Print(len(fruitSlice))

	fmt.Println(fruitSlice[1:2])

	alternateFruitSlice := []string{} //completely empty slice with no slice length or capacity
	alternateFruitSlice = append(alternateFruitSlice, "Mango")
	alternateFruitSlice = append(alternateFruitSlice, "Strawberries")
	alternateFruitSlice = append(alternateFruitSlice, "Sugarcane")

	fmt.Println(alternateFruitSlice)
	fmt.Println("Length: ", len(alternateFruitSlice))

}