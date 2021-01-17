package main

// struct is similar to classes

import (
	"fmt" 
	"strconv")

// Define person struct

type Person struct {
	/* firstName string
	lastName string
	city string
	gender string
	age int */

	// another way to define the struct - short form
	// same datatype in the same line
	
	firstName, lastName, city, gender string
	age int
}

// Value receiver method - used for calculations where data of struct is not changed
// Greeting method 
// basically we created a zero argument method and attached it to the Person struct with alias p
// In Java, we would have simply placed the method inside the class
func (p Person) greet() string{
	return "Hello, my name is "+p.firstName+" "+p.lastName + " and I am "+strconv.Itoa(p.age) //age is an int, so we can't simply put int there where return type is string
}

// pointer receiver
// we want to change the values inside the struct
func (p *Person) hasBirthday(){
	p.age++
}

// gets married (pointer reciever)
func (p *Person) getMarried(spouseLastName string){
	if(p.gender == "M"){
		return
	}else{
		p.lastName = spouseLastName
	}
}

func main(){
	// init person using struct

	person1 := Person{
		firstName:"Samantha",
		lastName: "Royce",
		city:"LA",
		gender:"F",
		age:38}

		person2 := Person{
			firstName:"Bob",
			lastName: "Johnson",
			city:"NY",
			gender:"M",
			age:34}

	//Alternative way to use struct without property name
	//person2:= Person{"Samantha","Royce","LA","F",38} // order should be same 
	/* fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1) */
	person1.hasBirthday()
	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())
	person2.getMarried("Thomson")
	fmt.Println(person2.greet())
	

}