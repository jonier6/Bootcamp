package main

import (
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("tasklist")

	err := cmd.Run()

	if err != nil {
		log.Fatalf("error executing command: %s", err)
	}

}
