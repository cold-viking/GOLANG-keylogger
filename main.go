package main

import (
	"fmt"
	//"os"

	hook "github.com/robotn/gohook"
)

func main() {
	fmt.Println("Start Program")
	//file, err := os.Open("log.txt")
	//if err != nil {
	//}
	//defer file.Close()

	events := hook.Start()

	for ev := range events {
		if ev.Kind == hook.KeyHold {
			if ev.Rawcode == 8 {
				fmt.Println("Backspace")
			} else if ev.Rawcode == 32 {
				fmt.Println("Space")
			} else if ev.Rawcode == 13 {
				fmt.Println("Enter")
			} else {
				fmt.Printf(string(ev.Keychar) + "\n")
			}
		}
	}
}
