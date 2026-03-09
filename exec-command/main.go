package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {

	cmd := exec.Command("tasklist")

	var buffer bytes.Buffer

	mw := io.MultiWriter(os.Stdout, &buffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	err := cmd.Run()

	if err != nil {
		log.Fatalf("error executing command: %s", err)
	}
	fmt.Println("\nCaptured output:")
	fmt.Println(buffer.String())

}
