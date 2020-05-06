package terminal

import (
	"fmt"
	"os"
	"runtime"
	"strconv"

	goTerminal "golang.org/x/crypto/ssh/terminal"
)

// Terminal -
type Terminal struct {
	file        *os.File
	theTerminal *goTerminal.Terminal
}

// NewTerminal -
func NewTerminal(file *os.File) *Terminal {
	return &Terminal{
		file:        file,
		theTerminal: goTerminal.NewTerminal(file, ""),
	}
}

// WriteString -
func (thisRef Terminal) WriteString(value string) {
	thisRef.theTerminal.Write([]byte(value))
}

// WriteBytes -
func (thisRef Terminal) WriteBytes(value []byte) {
	thisRef.theTerminal.Write(value)
}

// HideCursor -
func (thisRef Terminal) HideCursor() {
	if IsFileATerminal(thisRef.file) && runtime.GOOS != "windows" {
		thisRef.WriteString("\033[?25l")
	}
}

// ShowCursor -
func (thisRef Terminal) ShowCursor() {
	if IsFileATerminal(thisRef.file) && runtime.GOOS != "windows" {
		thisRef.WriteString("\033[?25h")
	}
}

// MoveCursorLineStart -
func (thisRef Terminal) MoveCursorLineStart() {
	thisRef.WriteString("\r")
}

// ClearLine -
func (thisRef Terminal) ClearLine() {
	width, _, _ := GetTerminalSize(thisRef.file)
	spaces := fmt.Sprintf("%"+strconv.Itoa(width)+"s", " ")

	thisRef.MoveCursorLineStart()
	thisRef.WriteString(spaces)
	thisRef.MoveCursorLineStart()
}

// ReadLine -
func (thisRef Terminal) ReadLine() (string, error) {
	return thisRef.theTerminal.ReadLine()
}

// ReadPassword -
func (thisRef Terminal) ReadPassword() (string, error) {
	return thisRef.theTerminal.ReadPassword("")
}
