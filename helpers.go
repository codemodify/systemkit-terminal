package terminal

import (
	"os"

	goTerminal "golang.org/x/crypto/ssh/terminal"
)

var goTerminalState *goTerminal.State

// IsFileDescriptorATerminal -
func IsFileDescriptorATerminal(fd uintptr) bool {
	return goTerminal.IsTerminal(int(fd))
}

// IsFileDescriptorATerminal2 -
func IsFileDescriptorATerminal2(fd int) bool {
	return goTerminal.IsTerminal(fd)
}

// IsFileATerminal -
func IsFileATerminal(file *os.File) bool {
	return goTerminal.IsTerminal(int(file.Fd()))
}

// GetTerminalSize -
func GetTerminalSize(file *os.File) (int, int, error) {
	return goTerminal.GetSize(int(file.Fd()))
}

// ReadPassword -
func ReadPassword(file *os.File) ([]byte, error) {
	return goTerminal.ReadPassword(int(file.Fd()))
}

// SaveState -
func SaveState(file *os.File) error {
	state, err := goTerminal.GetState(int(file.Fd()))
	if err != nil {
		return err
	}

	goTerminalState = state
	return nil
}

// RestoreState -
func RestoreState(file *os.File) error {
	if goTerminalState != nil {
		err := goTerminal.Restore(int(file.Fd()), goTerminalState)
		if err != nil {
			return err
		}
	}

	return nil
}
