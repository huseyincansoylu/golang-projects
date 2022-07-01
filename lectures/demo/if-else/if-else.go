package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
    quiz1, quiz2, quiz3 := 9,8,7

	if quiz1 > quiz2 {
		fmt.Println("quiz1 bigger than quiz2")
	}else if quiz1 < quiz2 {
		fmt.Println("quiz2 bigger than quiz1")
	} else {
		fmt.Println("quiz1 = quiz2")
	}

	if average(quiz1, quiz2 , quiz3) > 7 {
		fmt.Println("thats acceptable")
	} else {
		fmt.Println("no sorry")
	}
}
