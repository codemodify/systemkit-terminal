package terminal

// CursorShow -
func (thisRef Terminal) CursorShow() {
	thisRef.cursorShow()
}

// CursorHide -
func (thisRef Terminal) CursorHide() {
	thisRef.cursorHide()
}

// CursorMoveToLineStart -
func (thisRef Terminal) CursorMoveToLineStart() {
	thisRef.WriteString("\r")
}

// CursorMoveToXY -
func (thisRef Terminal) CursorMoveToXY(x int, y int) {
	thisRef.cursorMoveToXY(x, y)
}

// CursorMoveToX -
func (thisRef Terminal) CursorMoveToX(x int) {
	thisRef.cursorMoveToX(x)
}

// CursorMoveNLinesUp -
func (thisRef Terminal) CursorMoveNLinesUp(n int) {
	thisRef.cursorMoveNLinesUp(n)
}

// CursorMoveNLinesUpPositionStart -
func (thisRef Terminal) CursorMoveNLinesUpPositionStart(n int) {
	thisRef.cursorMoveNLinesUpPositionStart(n)
}

// CursorMoveNLinesDown -
func (thisRef Terminal) CursorMoveNLinesDown(n int) {
	thisRef.cursorMoveNLinesDown(n)
}

// CursorMoveNLinesDownPositionStart -
func (thisRef Terminal) CursorMoveNLinesDownPositionStart(n int) {
	thisRef.cursorMoveNLinesDownPositionStart(n)
}

// CursorMoveNColumnsLeft -
func (thisRef Terminal) CursorMoveNColumnsLeft(n int) {
	thisRef.cursorMoveNColumnsLeft(n)
}

// CursorMoveNColumnsRight -
func (thisRef Terminal) CursorMoveNColumnsRight(n int) {
	thisRef.cursorMoveNColumnsRight(n)
}

// CursorClearLineRight -
func (thisRef Terminal) CursorClearLineRight() {
	thisRef.cursorClearLineRight()
}

// CursorClearLineLeft -
func (thisRef Terminal) CursorClearLineLeft() {
	thisRef.cursorClearLineLeft()
}

// CursorClearLine -
func (thisRef Terminal) CursorClearLine() {
	thisRef.cursorClearLine()
}

// CursorClearScreenDown -
func (thisRef Terminal) CursorClearScreenDown() {
	thisRef.cursorClearScreenDown()
}

// CursorClearScreenUp -
func (thisRef Terminal) CursorClearScreenUp() {
	thisRef.cursorClearScreenUp()
}

// CursorPositionSave -
func (thisRef *Terminal) CursorPositionSave() {
	thisRef.cursorPositionSave()
}

// CursorPositionRestore -
func (thisRef Terminal) CursorPositionRestore() {
	thisRef.cursorPositionRestore()
}

// CursorPositionQuery -
func (thisRef Terminal) CursorPositionQuery() (int, int) {
	return thisRef.cursorPositionQuery()
}
