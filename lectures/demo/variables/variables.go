package main

import "fmt"

func main() {
	var myName = "Huseyin"
	fmt.Println("my name is",myName)

	var name string = "Karen"
	fmt.Println("name =", name)

    username := "admin"
	fmt.Println("username=", username)

    var sum int
	fmt.Println("The sum is=", sum) //default 0

	part1, other := 1, 5
	fmt.Println("part1=", part1, "other is=", other)
     
	//other will be  5
	
	part2, other := 2, 0
	fmt.Println("part2=", part2, "other is=", other)

	//other will be  0


	sum = part1 + part2
	fmt.Println("sum=", sum)

	var(
		lessonName = "Variables"
		lessonType = "Demo"
	)

	fmt.Println(lessonName)
	fmt.Println(lessonType)


	word1, word2, _ := "hello", "world", "1"

	fmt.Println(word1, word2)
}
