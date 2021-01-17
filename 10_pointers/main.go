package main

import "fmt"

func main(){
	a := 5

	// to store address of value in variable, use &
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n",b)

	// Use * to read val from address, basically to get value from an addrees use *
	fmt.Println(*b)
	fmt.Println(*&a)
	

	//change value with pointer
	*b = 10
	fmt.Println(a) // prints 10 because b points to a's address. 
	//changing value stored in b essentially means changing value stored in the memory location of a. So, a will also reflect that change

	// everything in GO is passed by value
}