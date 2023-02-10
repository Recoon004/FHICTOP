package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	fmt.Println("Welkom bij Fonteyn Vakantieparken!")
	if hour > 23 {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")
	} else if hour > 18 {
		fmt.Println("goedenavond!")
	} else if hour > 13 {
		fmt.Println("goedemidag!")
	} else if hour > 7 {
		fmt.Println("goedemorgen!")
	} else {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")
	}
}
