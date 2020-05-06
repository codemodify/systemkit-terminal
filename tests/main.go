package main

import (
	"fmt"
	"os"
	"time"

	terminal "github.com/codemodify/systemkit-terminal"
)

func main() {

	timeToSleep := 1 * time.Second

	// Test 1
	term := terminal.NewTerminal(os.Stdout)

	term.ShowCursor()
	term.MoveCursorLineStart()
	term.WriteString("cursor shown")
	time.Sleep(timeToSleep)

	term.HideCursor()
	term.MoveCursorLineStart()
	term.WriteString("cursor hidden")
	time.Sleep(timeToSleep)

	term.ShowCursor()
	term.MoveCursorLineStart()
	term.WriteString("cursor shown again")
	time.Sleep(timeToSleep)

	// Test
	fmt.Println()
	fmt.Println()

	term.HideCursor()
	time.Sleep(timeToSleep)
	fmt.Println("aaaaaaaaaaaaaaa")
	fmt.Println("bbbbbbbbbbbbbbb")

	time.Sleep(timeToSleep)
	fmt.Print("ccccccccccccccc")
	time.Sleep(timeToSleep)
	term.ClearLine()
	fmt.Print("ddddddd")
	time.Sleep(timeToSleep)

	fmt.Println()
	term.ShowCursor()
}
