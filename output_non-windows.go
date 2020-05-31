// +build !windows

package terminal

func (thisRef Terminal) clearScreen() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.WriteString(escape("[2J"))
}
