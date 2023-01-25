package main

import (
	"fmt"
	"projectSOLOMON/cerebrum"
	"projectSOLOMON/ears"
	"projectSOLOMON/librarian"
	"projectSOLOMON/tongue"
)

func main() {
	fmt.Println("Main Fucntion Start")

	processInput()

	fmt.Println("Main Fucntion End")
}

func processInput() {
	fmt.Println("Process Input Fucntion Start")

	ears.Listen()
	cerebrum.Cerebrate()
	librarian.Scower()
	tongue.Speak()

	fmt.Println("Process Input Fucntion Start")
}
