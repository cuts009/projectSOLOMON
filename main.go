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

	listenSpeak()

	fmt.Println("Main Fucntion End")
}

func listenSpeak() {
	fmt.Println("Process-Respond Fucntion Start")

	ears.Listen()
	cerebrum.Cerebrate()
	skills.Initiate()
	tongue.Speak()

	fmt.Println("Process-Respond Fucntion End")
}
