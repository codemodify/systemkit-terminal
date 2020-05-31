// +build windows

package terminal

import (
	"strings"
	"syscall"
	"unsafe"
)

func (thisRef Terminal) cursorShow() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	set_console_cursor_info(
		syscall.Handle(thisRef.file.Fd()),
		&console_cursor_info{
			size:    100,
			visible: 1,
		},
	)
}

func (thisRef Terminal) cursorHide() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	set_console_cursor_info(
		syscall.Handle(thisRef.file.Fd()),
		&console_cursor_info{
			size:    100,
			visible: 0,
		},
	)
}

func (thisRef Terminal) cursorMoveToXY(x int, y int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	set_console_cursor_position(
		syscall.Handle(thisRef.file.Fd()),
		coord{short(x), short(y)},
	)
}

func (thisRef Terminal) cursorMoveToX(x int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	currentCoord.x = short(x)

	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorMoveNLinesUp(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	currentCoord.y = currentCoord.y - short(n)

	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorMoveNLinesUpPositionStart(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	currentCoord.x = 0
	currentCoord.y = currentCoord.y - short(n)

	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorMoveNLinesDown(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	currentCoord.y = currentCoord.y + short(n)

	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorMoveNLinesDownPositionStart(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	currentCoord.x = 0
	currentCoord.y = currentCoord.y + short(n)

	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorMoveNColumnsLeft(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	currentCoord.x = currentCoord.x - short(n)

	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorMoveNColumnsRight(n int) {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	currentCoord.x = currentCoord.x + short(n)

	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorClearLineRight() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	// get current cursor X,Y and window size
	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	windowSize := get_window_size(syscall.Handle(thisRef.file.Fd()))

	// write spaces until end of the line
	manySpaces := strings.Repeat(" ", int(windowSize.x-currentCoord.x))
	thisRef.WriteString(manySpaces)

	// restore cursor X,Y
	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorClearLineLeft() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	// get current cursor X,Y and window size
	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	// windowSize := get_window_size(syscall.Handle(thisRef.file.Fd()))

	// write spaces until end of the line
	manySpaces := strings.Repeat(" ", int(currentCoord.x))
	thisRef.cursorMoveToXY(0, int(currentCoord.y))
	thisRef.WriteString(manySpaces)

	// restore cursor X,Y
	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorClearLine() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	// get current cursor X,Y and window size
	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	windowSize := get_window_size(syscall.Handle(thisRef.file.Fd()))

	// write spaces until end of the line
	manySpaces := strings.Repeat(" ", int(windowSize.x))
	thisRef.cursorMoveToXY(0, int(currentCoord.y))
	thisRef.WriteString(manySpaces)

	// restore cursor X,Y
	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorClearScreenDown() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	// get current cursor X,Y and window size
	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	windowSize := get_window_size(syscall.Handle(thisRef.file.Fd()))

	// write spaces until end of the line
	line := strings.Repeat(" ", int(windowSize.x-1)) + "\n"
	lines := strings.Repeat(line, int(windowSize.y-currentCoord.y))

	thisRef.cursorMoveToXY(0, int(currentCoord.y))
	thisRef.WriteString(lines)

	// restore cursor X,Y
	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef Terminal) cursorClearScreenUp() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	// get current cursor X,Y and window size
	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	windowSize := get_window_size(syscall.Handle(thisRef.file.Fd()))

	// write spaces until end of the line
	line := strings.Repeat(" ", int(windowSize.x-1)) + "\n"
	lines := strings.Repeat(line, int(currentCoord.y))

	thisRef.cursorMoveToXY(0, 0)
	thisRef.WriteString(lines)

	// restore cursor X,Y
	thisRef.cursorMoveToXY(int(currentCoord.x), int(currentCoord.y))
}

func (thisRef *Terminal) cursorPositionSave() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.cursortCurrentX, thisRef.cursortCurrentY = thisRef.cursorPositionQuery()
}

func (thisRef Terminal) cursorPositionRestore() {
	if !IsFileATerminal(thisRef.file) {
		return
	}

	thisRef.cursorMoveToXY(thisRef.cursortCurrentX, thisRef.cursortCurrentY)
}

func (thisRef Terminal) cursorPositionQuery() (int, int) {
	if !IsFileATerminal(thisRef.file) {
		return 0, 0
	}

	currentCoord := get_cursor_position(syscall.Handle(thisRef.file.Fd()))
	return int(currentCoord.x), int(currentCoord.y)
}

// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// WINDOWS specific APIs
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~

type (
	dword uint32
	short int16
	word  uint16
	wchar uint16

	small_rect struct {
		left   short
		top    short
		right  short
		bottom short
	}

	console_cursor_info struct {
		size    dword
		visible int32
	}

	console_screen_buffer_info struct {
		size                coord
		cursor_position     coord
		attributes          word
		window              small_rect
		maximum_window_size coord
	}

	coord struct {
		x short
		y short
	}
)

var (
	kernel32                            = syscall.NewLazyDLL("kernel32.dll")
	proc_set_console_cursor_position    = kernel32.NewProc("SetConsoleCursorPosition")
	proc_set_console_cursor_info        = kernel32.NewProc("SetConsoleCursorInfo")
	proc_get_console_screen_buffer_info = kernel32.NewProc("GetConsoleScreenBufferInfo")
	proc_fill_console_output_attribute  = kernel32.NewProc("FillConsoleOutputAttribute")
	proc_fill_console_output_character  = kernel32.NewProc("FillConsoleOutputCharacterW")
)

func (this coord) uintptr() uintptr {
	return uintptr(*(*int32)(unsafe.Pointer(&this)))
}

func set_console_cursor_info(h syscall.Handle, info *console_cursor_info) (err error) {
	r0, _, e1 := syscall.Syscall(proc_set_console_cursor_info.Addr(),
		2, uintptr(h), uintptr(unsafe.Pointer(info)), 0)
	if int(r0) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func set_console_cursor_position(h syscall.Handle, pos coord) (err error) {
	r0, _, e1 := syscall.Syscall(proc_set_console_cursor_position.Addr(),
		2, uintptr(h), pos.uintptr(), 0)
	if int(r0) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func get_cursor_position(out syscall.Handle) coord {
	var tmp_info console_screen_buffer_info
	err := get_console_screen_buffer_info(out, &tmp_info)
	if err != nil {
		return coord{}
	}
	return tmp_info.cursor_position
}

func get_window_size(out syscall.Handle) coord {
	var tmp_info console_screen_buffer_info
	err := get_console_screen_buffer_info(out, &tmp_info)
	if err != nil {
		return coord{}
	}
	return tmp_info.maximum_window_size
}

// func get_window_info(out syscall.Handle) console_screen_buffer_info {
// 	var tmp_info console_screen_buffer_info
// 	err := get_console_screen_buffer_info(out, &tmp_info)
// 	if err != nil {
// 		return console_screen_buffer_info{}
// 	}
// 	return tmp_info
// }

func get_console_screen_buffer_info(h syscall.Handle, info *console_screen_buffer_info) (err error) {
	r0, _, e1 := syscall.Syscall(proc_get_console_screen_buffer_info.Addr(),
		2, uintptr(h), uintptr(unsafe.Pointer(info)), 0)
	if int(r0) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func fill_console_output_attribute(h syscall.Handle, attr word, n int) (err error) {
	tmp_coord := coord{0, 0}
	var tmp_arg dword
	r0, _, e1 := syscall.Syscall6(proc_fill_console_output_attribute.Addr(),
		5, uintptr(h), uintptr(attr), uintptr(n), tmp_coord.uintptr(),
		uintptr(unsafe.Pointer(&tmp_arg)), 0)
	if int(r0) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func fill_console_output_character(h syscall.Handle, char wchar, n int) (err error) {
	tmp_coord := coord{0, 0}
	var tmp_arg dword
	r0, _, e1 := syscall.Syscall6(proc_fill_console_output_character.Addr(),
		5, uintptr(h), uintptr(char), uintptr(n), tmp_coord.uintptr(),
		uintptr(unsafe.Pointer(&tmp_arg)), 0)
	if int(r0) == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
