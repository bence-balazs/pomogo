package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

var counter int

func main() {

	// set cli flags
	tts := flag.String("tts", "endofpomodo", "tts message")
	timer := flag.Int("time", 5, "pomodo time")
	flag.Parse()

	// check if there is any args given
	if flag.NFlag() < 1 {
		fmt.Println("usage: ./pomogo -time=60 -tts=hello")
		os.Exit(0)
	}

	// start program logic
	counter = 0
	for {
		fmt.Printf("Time remaining: %v min.\n", *timer)
		remainingTimeCounter(*timer)
		sendSoundNotify(*tts)
		err := sendPopupNotify()
		counter++
		if err != nil {
			fmt.Printf("You have completed %v pomodo(s).\n", counter)
			os.Exit(0)
		}
	}
}

// sendSoundNotify is sending a builtin linux tts notification
func sendSoundNotify(message string) {
	_, err := exec.Command("spd-say", message).Output()
	if err != nil {
		log.Fatal(err)
	}
}

// sendPopupNotify is sending a builtin linux popup notification
func sendPopupNotify() error {
	_, err := exec.Command("zenity", "--question").Output()
	return err
}

// remainingTimeCounter counts down every minute and prints out the remaining time
func remainingTimeCounter(timer int) {
	for range time.Tick(1 * time.Minute) {
		if timer == 1 {
			fmt.Println("End of this pomodo session!")
			break
		}
		timer--
		fmt.Printf("Time remaining: %v min.\n", timer)
	}
}
