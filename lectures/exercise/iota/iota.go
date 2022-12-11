//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	Add = iota
	Subtract
	Multiply
	Divide 
)

type Operation int

func (op Operation ) calculate(left , right int) int { 
	switch op {
		case Add:
			return left + right
		case Subtract:
			return left - right
		case Multiply:
			return left * right
		case Divide: 
		    return left / right 
	}
	panic("unhandled operations" )
}

func main() {
    add := Operation(Add)
    sub := Operation(Subtract)
    mul := Operation(Multiply)
    div := Operation(Divide)

	fmt.Println(add.calculate(2, 2)) // = 4

	fmt.Println(sub.calculate(10, 3)) // = 7

	fmt.Println(mul.calculate(3, 3)) // = 9

	fmt.Println(div.calculate(100, 2)) // = 50
}
