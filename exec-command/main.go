package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("tasklist")

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	if err != nil {
		log.Fatalf("error executing command: %s", err)
	}
	fmt.Println("STDOUT:")
	fmt.Println(stdout.String())

	fmt.Println("STDERR:")
	fmt.Println(stderr.String())

}
