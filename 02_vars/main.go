package main

import "fmt"

/* // global variable works like it should
var name = "Ronet" */

func main(){

	// Using vars
	name:= "Ronet"	// shorthand notation, can only be used inside a function

	// we always need to use the variable that being declared in go, otherwise error NOT warning
	var age int32 = 23
	const isCool = true;
	size := 1.3; //float64 default

	// multiple assignments
	name, email := "Ronet Swaminathan", "ronet_swaminathan@ymail.com"

	fmt.Println(name, age, isCool, email);
	fmt.Printf("%T\n", size); // find the type of a variable. %T is the verb

}