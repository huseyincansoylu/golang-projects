package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(a, b int) int {
	return a + b
}

func greet() {
	fmt.Println("Hello from greet func!")
}

func main() {
    greet()

	dozen := double(6)

	fmt.Println("A dozen is = " , dozen)

	bakersDozen := add(dozen , 1)

	fmt.Println("A bakersDozen is = " , bakersDozen)

	anotherBakersDozen := add(double(6), 9)
	fmt.Println("A anotherBakersDozen is = " , anotherBakersDozen)
}
