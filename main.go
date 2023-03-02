package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type configuration struct {

		port int
		pswd string
		server string
		Username string `json:"Username"`
		Databasename string `json:"Databasename"`

}
var config configuration error
file, _ := os.Open("DB.json")
json := json.NewDecoder(file)
json.Decode(&config)

func KentekeninGeven() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type je kenteken: ")
	scanner.Scan()
	input := scanner.Text()
	if input == "Ka-as-01" {
		GetTime()
	} else {
		fmt.Println("geen overeen koment kenteken")
	}
	fmt.Println("je kentaken is:", input)
}
func GetTime() {
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
