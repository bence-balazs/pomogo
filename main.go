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
	tts := flag.String("tts", "pomodo", "tts message")
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
		time.Sleep(time.Second * time.Duration(*timer))
		sendSoundNotify(*tts)
		err := sendPopupNotify()
		counter++
		if err != nil {
			fmt.Printf("You have completed %v pomodo(s).\n", counter)
			os.Exit(0)
		}
	}
}

func sendSoundNotify(message string) {
	_, err := exec.Command("spd-say", message).Output()
	if err != nil {
		log.Fatal(err)
	}
}

func sendPopupNotify() error {
	_, err := exec.Command("zenity", "--question").Output()
	return err
}
