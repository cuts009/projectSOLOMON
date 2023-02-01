// Copyright (C) 2023, Cutter Love (cuts009dev@proton.me)
// Licensed under the GPL v3.0, you may obtain a copy of this license at:
// https://www.gnu.org/licenses/gpl-3.0-standalone.html

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

	startUp()

	listenSpeak()

	fmt.Println("Main Fucntion End")
}

func startUp() {
	fmt.Println("Start-Up Function Start")

	fmt.Println("Hello, I am SOLOMON. How can I help you today?")

	idle()

	fmt.Println("Start-Up Function End")

}

func idle() {
	fmt.Println("Idle Function Start")

	fmt.Println("Idle Function End")
}

func listenSpeak() {
	fmt.Println("Process-Respond Fucntion Start")

	ears.Listen()
	cerebrum.Cerebrate()
	skills.Initiate()
	tongue.Speak()

	fmt.Println("Process-Respond Fucntion End")
}
