package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start Program")

	if err := startEventWriter(); err != nil {
		fmt.Fprintf(os.Stderr, "Error starting event log writer: %v\n", err)
		os.Exit(1)
	}
}
