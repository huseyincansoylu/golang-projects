package main

import "fmt"

type Passenger struct {
	 Name  string
	 TicketNumber int
	 Boarded bool
}

type Bus struct {
    FrontSeat Passenger 
}

func main() {
     can := Passenger{"Can", 1, false}
	 fmt.Println(can)

	 var (
		buse = Passenger{Name: "Buse", TicketNumber: 2}
		ahmet = Passenger{Name: "Ahmet", TicketNumber: 3} 
	 )
	  
	 fmt.Println(buse, ahmet)
	 
	 var heidi  Passenger
	 heidi.Name = "Heidi"
	 heidi.TicketNumber = 4 
	 fmt.Println(heidi) 

	 can.Boarded = true
	 ahmet.Boarded = true
	  
	 if ahmet.Boarded {
		fmt.Println("ahmet boarded")
	 }

	 if can.Boarded {
        fmt.Println("ahmet boarded")
	 }  

	 heidi.Boarded = true

	 bus := Bus{heidi}

	 fmt.Println(bus) 
	 fmt.Println(bus.FrontSeat.Name)
}
