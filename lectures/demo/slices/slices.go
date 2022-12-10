package main

import "fmt"


func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("----", title, "----")
	for i := 0; i < len(slice); i++ {
		element := slice[i]

		fmt.Println(element)
	}
}

func main() {
     route := []string{"grocery", "departmant store", "salon"}

	 printSlice("1" , route)

	 route = append(route, "theatre")

	 printSlice("2" , route)
	
	 fmt.Println()

	 fmt.Println(route[0], "is visited")
	 fmt.Println(route[1], "is visited")


	 route = route[2:]
	 fmt.Println("remaining routes", route)
}
