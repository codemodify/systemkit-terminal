// +build windows

package terminal

import "syscall"

func (thisRef Terminal) clearScreen() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	windowSize := get_window_size(syscall.Handle(thisRef.file.Fd()))

	area := int(windowSize.x) * int(windowSize.y)
	fill_console_output_attribute(syscall.Handle(thisRef.file.Fd()), 0, area)
	fill_console_output_character(syscall.Handle(thisRef.file.Fd()), 0, area)
}
