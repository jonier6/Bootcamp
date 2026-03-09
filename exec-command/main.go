package main

import (
	"log"
	"os/exec"
	"os"
)

func main() {

	cmd := exec.Command("tasklist")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	err := cmd.Run()

	if err != nil {
		log.Fatalf("error executing command: %s", err)
	}

}
