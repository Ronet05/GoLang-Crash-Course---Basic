package main

import "fmt"

func main(){
	// define maps, basically dict
	emails := make(map[string]map[int]string)
	
	//assign key value
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Bob"] = "bob1@gmail.com"
	emails["Male"] = make(map[int]string)
	emails["Female"] = make(map[int]string)

	emails["Male"][0] = "bob@gmail.com"
	emails["Male"][1] = "someoneelse@gmail.com"
	emails["Female"][0] = "sharon@gmail.com"

	// delete from map
	delete(emails["Male"], 1)

	// Declare another way without make
	names := map[int]string{} 
	names[0] = "Ronet Swaminathan"
	fmt.Println(names)

	fmt.Println(emails)

}