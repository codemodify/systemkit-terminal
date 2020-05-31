package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	terminal "github.com/codemodify/systemkit-terminal"
)

const sleepTime = 1 * time.Second
const startX = 5
const startY = 5

type demo func(*terminal.Terminal)

func getFunctionName() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())

	return f.Name()
}

func main() {
	theTerminal := terminal.NewTerminal(os.Stdout)

	demos := []demo{
		test_CursorHide,
		test_CursorShow,
		test_CursorMoveToLineStart,
		test_CursorMoveToXY,
		test_CursorMoveToX,
		test_CursorMoveNLinesUp,
		test_CursorMoveNLinesUpPositionStart,
		test_CursorMoveNLinesDown,
		test_CursorMoveNLinesDownPositionStart,
		test_CursorMoveNColumnsLeft,
		test_CursorMoveNColumnsRight,
		test_CursorClearLineRight,
		test_CursorClearLineLeft,
		test_CursorClearLine,
		test_CursorClearScreenDown,
		test_CursorClearScreenUp,
		test_CursorPositionSave_CursorPositionRestore,
		test_CursorPositionQuery,

		test_SomethingRandom,
	}

	for _, demo := range demos {
		demo(theTerminal)
		time.Sleep(sleepTime)
	}

	fmt.Println()
	fmt.Println()
}

func test_CursorHide(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will be hidden in %s seconds", sleepTime)

	time.Sleep(sleepTime)
	t.CursorHide()
}

func test_CursorShow(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will be shown in %s", sleepTime)

	time.Sleep(sleepTime)
	t.CursorShow()
}

func test_CursorMoveToLineStart(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move to the start of the line in %s and print X", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveToLineStart()
	t.WriteString("X")
}

func test_CursorMoveToXY(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move to position 15 x 15 in %s and print HELLO", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveToXY(15, 15)
	t.WriteString("HELLO")
}

func test_CursorMoveToX(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move to position %d in %s and print capital the same", startX, sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveToX(startX)
	t.WriteString("CURSOR WILL MOVE")
}

func test_CursorMoveNLinesUp(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move 1 line up in %s and print SOMETHING", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveNLinesUp(1)
	t.WriteString("SOMETHING")
}

func test_CursorMoveNLinesUpPositionStart(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move 1 line up and start of the line up in %s and print SOMETHING", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveNLinesUpPositionStart(1)
	t.WriteString("SOMETHING")
}

func test_CursorMoveNLinesDown(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move 1 line down in %s and print SOMETHING", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveNLinesDown(1)
	t.WriteString("SOMETHING")
}

func test_CursorMoveNLinesDownPositionStart(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move 1 line down and start of the line up in %s and print SOMETHING", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveNLinesDownPositionStart(1)
	t.WriteString("SOMETHING")
}

func test_CursorMoveNColumnsLeft(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move to 57 positions to left in %s and print capital MOVE", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveNColumnsLeft(57)
	t.WriteString("MOVE")
}

func test_CursorMoveNColumnsRight(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("cursor will move to 1 positions to right in %s and print CAPITAL", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveNColumnsRight(1)
	t.WriteString("CAPITAL")
}

func test_CursorClearLineRight(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("line will clear to the left of THIS in %s", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveNColumnsLeft(5)
	time.Sleep(sleepTime)
	t.CursorClearLineRight()
}

func test_CursorClearLineLeft(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("line will clear to the left of THIS in %s", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveNColumnsLeft(11)
	time.Sleep(sleepTime)
	t.CursorClearLineLeft()
}

func test_CursorClearLine(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("line will clear in %s", sleepTime)

	time.Sleep(sleepTime)
	t.CursorClearLine()
}

func test_CursorClearScreenDown(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("screen below will clear in %s", sleepTime)

	t.CursorMoveToXY(startX, startY+5)
	t.WriteStringf("SOME NOISE TO CLEAR")

	t.CursorMoveToXY(startX, startY+3)

	time.Sleep(sleepTime)
	t.CursorClearScreenDown()
}

func test_CursorClearScreenUp(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("screen above will clear in %s", sleepTime)

	t.CursorMoveToXY(startX, startY+1)

	time.Sleep(sleepTime)
	t.CursorClearScreenUp()
}

func test_CursorPositionSave_CursorPositionRestore(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("saved cursors position, restore will be here->")

	t.CursorPositionSave()

	t.CursorMoveToXY(startX, startY+3)
	t.WriteStringf("going to cursors 15 x 15 in %s and print SOMETHING", sleepTime)

	time.Sleep(sleepTime)
	t.CursorMoveToXY(15, 15)
	t.WriteString("SOMETHING")

	time.Sleep(sleepTime)
	t.CursorMoveToXY(startX, startY+4)
	t.WriteStringf("restoring cursors position in %s and print RESTORED", sleepTime)

	time.Sleep(sleepTime)
	t.CursorPositionRestore()
	t.WriteString("RESTORED")
}

func test_CursorPositionQuery(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	t.CursorMoveToXY(startX, startY+2)
	t.WriteStringf("querying cursor position in %s", sleepTime)

	time.Sleep(sleepTime)
	col, line := t.CursorPositionQuery()
	t.CursorMoveToXY(startX, startY+3)
	t.WriteStringf("cursor position: %d x %d", col, line)
}

func test_SomethingRandom(t *terminal.Terminal) {
	t.ClearScreen()

	t.CursorMoveToXY(startX, startY)
	t.WriteString("DEMO: " + getFunctionName())

	for i := 0; i < 20; i++ {
		t.CursorMoveToXY(i, i)
		fmt.Print("X")
	}
}
