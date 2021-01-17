package main

import "fmt"

func main() {
	x := 5
	y := 10

	if x <= y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//  else if 
	color := "Blue"
	if(color == "Red"){
		fmt.Println("Color is Red")
	}else if (color == "Blue"){
		fmt.Println("Color is Blue")
	}else{
		fmt.Println("Color is not blue or red")
	}

	// Switch
	switch color{
	case "Red":
		fmt.Println("Color is red")
	case "Blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not blue or red")
	}
}
