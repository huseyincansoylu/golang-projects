package main

import "fmt"

type Room struct {
	name string
	cleaned bool
}

func check(rooms [4]Room) {
    for i := 0; i < len(rooms); i++ {
         room := rooms[i]
		 if room.cleaned {
            fmt.Println(room.name, "is clean")
		 } else {
			fmt.Println(room.name, "is not clean")
		 }
	}
} 

func main() {
    rooms := [...]Room{
		{name:  "salon", cleaned: true},
		{name:  "mutfak", cleaned: false},
		{name:  "banyo", cleaned: false},
		{name:  "yatak odasi", cleaned: true},
	}

	check(rooms) 
}
