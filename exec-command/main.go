package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("tasklist")

	var buffer bytes.Buffer

	cmd.Stdout = &buffer
	cmd.Stderr = &buffer

	err := cmd.Run()

	if err != nil {
		log.Fatalf("error executing command: %s", err)
	}
	fmt.Println("Combined output:")
	fmt.Println(buffer.String())

}
