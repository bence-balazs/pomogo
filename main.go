package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const message = "pomodo end"

var counter int

func main() {
	counter = 0
	for {
		time.Sleep(time.Second * 5)
		sendSoundNotify(message)
		err := sendPopupNotify()
		if err != nil {
			fmt.Printf("You have completed %v pomodo(s).\n", counter)
			os.Exit(0)
		}
		counter++
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
