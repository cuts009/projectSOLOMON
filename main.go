package main

import (
	"fmt"
	"projectSOLOMON/cerebrum"
	"projectSOLOMON/ears"
	"projectSOLOMON/skills"
	"projectSOLOMON/tongue"
)

func main() {
	fmt.Println("Main Fucntion Start")

	processRespond()

	fmt.Println("Main Fucntion End")
}

func processRespond() {
	fmt.Println("Process-Respond Fucntion Start")

	ears.Listen()
	cerebrum.Cerebrate()
	skills.Initiate()
	tongue.Speak()

	fmt.Println("Process-Respond Fucntion End")
}
