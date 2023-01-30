package main

import (
	"fmt"
	"projectSOLOMON/assistant"
	"projectSOLOMON/cerebrum"
	"projectSOLOMON/ears"
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
	assistant.Help()
	tongue.Speak()

	fmt.Println("Process-Respond Fucntion End")
}
