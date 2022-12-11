package main

import "fmt"

type Preparer interface {
     PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish() {
	fmt.Println("cook chicken")
}

func (s Salad) PrepareDish() {
	fmt.Println("chop  salad")
	fmt.Println("add  dressing ")
}

func prepareDishes(dishes []Preparer ) {
	 fmt.Println("prepare dishes ")
	 for i := 0; i < len(dishes); i++ {
		  dish := dishes[i]  
		  fmt.Println(dish)
		  dish.PrepareDish() 
	 }
	 fmt.Println( )
}
   
func main() {
     dishes := []Preparer{Chicken("Chicken"), Salad("saladdd ")}

	 prepareDishes(dishes)
}
