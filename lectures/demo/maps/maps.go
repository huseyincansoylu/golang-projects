package main

import "fmt"



func main() {
    shoppingList := make(map[string]int)

	shoppingList["eggs"] = 11
	shoppingList["milk"] = 5
	shoppingList["bread"] += 1 

	shoppingList["eggs"] += 9

	fmt.Println(shoppingList)

	delete(shoppingList, "milk")

	fmt.Println(shoppingList) 

	cereal, found := shoppingList["cereal"]

	if !found {
		fmt.Println(("no cereal"))
	} else {
		fmt.Println(cereal)
	}
  
	totalItems := 0

	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println(totalItems)
	 
}


