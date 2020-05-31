// +build !windows

package terminal

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	goTerminal "golang.org/x/crypto/ssh/terminal"
)

const escapeCode = "\x1b"

func escape(format string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", escapeCode, fmt.Sprintf(format, args...))
}

func (thisRef Terminal) cursorShow() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[?25h"))
}

func (thisRef Terminal) cursorHide() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[?25l"))
}

func (thisRef Terminal) cursorMoveToXY(x int, y int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[%d;%dH", y, x))
}

func (thisRef Terminal) cursorMoveToX(x int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[%dG", x))
}

func (thisRef Terminal) cursorMoveNLinesUp(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[%dA", n))
}

func (thisRef Terminal) cursorMoveNLinesUpPositionStart(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[%dF", n))
}

func (thisRef Terminal) cursorMoveNLinesDown(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[%dB", n))
}

func (thisRef Terminal) cursorMoveNLinesDownPositionStart(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[%dE", n))
}

func (thisRef Terminal) cursorMoveNColumnsLeft(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[%dD", n))
}

func (thisRef Terminal) cursorMoveNColumnsRight(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[%dC", n))
}

func (thisRef Terminal) cursorClearLineRight() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[0K"))
}

func (thisRef Terminal) cursorClearLineLeft() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[1K"))
}

func (thisRef Terminal) cursorClearLine() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[2K"))
}

func (thisRef Terminal) cursorClearScreenDown() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[0J"))
}

func (thisRef Terminal) cursorClearScreenUp() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[1J"))
}

func (thisRef Terminal) cursorPositionSave() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[s"))
}

func (thisRef Terminal) cursorPositionRestore() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[u"))
}

func (thisRef Terminal) cursorPositionQuery() (int, int) {
	if !IsFileATerminal(thisRef.file) {
		return 0, 0
	}

	// 1. make it raw, save state
	state, _ := goTerminal.MakeRaw(int(thisRef.file.Fd()))

	// 2. magic
	thisRef.WriteString(escape("[6n"))
	reader := bufio.NewReader(thisRef.file)
	data, _ := reader.ReadBytes('R')

	// 3. magic
	goTerminal.Restore(int(thisRef.file.Fd()), state)

	// 4. parse line and column
	dataAsString := string(data)
	dataAsString = strings.Replace(dataAsString, escapeCode, "", 1)
	dataAsString = strings.Replace(dataAsString, "[", "", 1)
	dataAsString = strings.Replace(dataAsString, "R", "", 1)
	colLineAsArray := strings.Split(dataAsString, ";")
	col := 0
	line := 0
	if len(colLineAsArray) > 0 {
		line, _ = strconv.Atoi(colLineAsArray[0])
	}
	if len(colLineAsArray) > 1 {
		col, _ = strconv.Atoi(colLineAsArray[1])
	}

	return col, line
}
