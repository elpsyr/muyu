package main

import (
	"fmt"
	"time"
)

func main() {
	frames := []string{
		`
 (\(\            
 ( -.-)  ∕(i) 
 o_(")(")
		`,
		`
 (\(\      功德+1
 ( -.-)  ㇀(i)  
 o_(")(")
		`,
	}

	clearScreen()

	for {
		for _, frame := range frames {
			moveCursorToTop()

			clearFrame()
			fmt.Print(frame)

			time.Sleep(300 * time.Millisecond)
		}
	}
}

func clearScreen() {
	fmt.Print("\033[2J")
}

func moveCursorToTop() {
	fmt.Print("\033[H")
}

func clearFrame() {
	fmt.Print("\033[2K")
	moveCursorToTop()
}
