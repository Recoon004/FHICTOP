package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	fmt.Println("Welkom bij Fonteyn Vakantieparken!")
	if hour == 7 {
		fmt.Println("goedemorgen!")
	}
	if hour == 13 {
		fmt.Println("goedemidag!")
	}
	if hour == 18 {
		fmt.Println("goedenavond!")
	}
	if hour == 23 {
		fmt.Println("Sorry, de parkeerplaats is ’s nachts gesloten")
	}
}
