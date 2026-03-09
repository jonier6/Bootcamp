package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("tasklist")

	output, err := cmd.Output()

	if err != nil {
		log.Fatalf("error executing command: %s", err)
	}
	fmt.Println("Command output:")
	fmt.Println(string(output))

}
