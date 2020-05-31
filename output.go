package terminal

import "fmt"

// WriteString -
func (thisRef Terminal) WriteString(value string) {
	thisRef.theTerminal.Write([]byte(value))
}

// WriteStringf -
func (thisRef Terminal) WriteStringf(format string, v ...interface{}) {
	thisRef.theTerminal.Write([]byte(fmt.Sprintf(format, v...)))
}

// WriteBytes -
func (thisRef Terminal) WriteBytes(value []byte) {
	thisRef.theTerminal.Write(value)
}

// ClearScreen -
func (thisRef Terminal) ClearScreen() {
	thisRef.clearScreen()
}
