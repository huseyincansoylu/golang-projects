package main

import "fmt"

func main() {
	slice := []string{"hello", "world", "!"}

	 for i, element := range slice{
       fmt.Println(i, element)
	   for _, letter := range element {
		fmt.Printf("char %q\n", letter)
	   }
	 }
}
