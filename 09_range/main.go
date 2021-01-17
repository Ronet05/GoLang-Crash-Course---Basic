package main

import "fmt"

func main() {
	ids := []int{33, 67, 34, 7, 34, 86}

	// loops through ids using range, i is for index and id is for each element
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// not using index - _ for avoiding to use index variable, because if used in range loop, then it HAS be used somewhere or else error
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// adding the ids together
	sum := 0
	for _, id := range ids {
		sum+= id
	}
	fmt.Println("The sum of ids = ", sum)

	//Range with map and add kv
	emails := map[string]string {"Bob":"bob@gmail.com", "Sharon":"shardon@gmail.com"}
	// index becomes the key, so it is same analogy
	for k,v := range emails{
		fmt.Printf("Name: %s, Email: %s\n", k,v)
	}

	// only the indexes or the keys
	for k := range emails{
		fmt.Printf("Name: %s\n", k)
	}

	// only the values
	for _,v := range emails{
		fmt.Printf("Email: %s\n", v)
	}

}
