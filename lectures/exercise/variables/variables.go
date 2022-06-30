//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	//Requirements:
//* Store your favorite color in a variable using the `var` keyword
     var myFavColor = "black"
	 fmt.Println("my fav color=", myFavColor)

//* Store your birth year and age (in years) in two variables using
//  compound assignment
    year, age := 1993, 29
	fmt.Println("my birth year=", year, "my age=",age)
    
//* Store your first & last initials in two variables using block assignment
       var(
		 firstInitial = "J"
		 lastInitial = "K"
	   )
       
	   fmt.Println(firstInitial, lastInitial)

//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier

    var ageInDays int

    ageInDays = 365 * age
    
	fmt.Println(ageInDays)


}
