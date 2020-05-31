package terminal

import (
	"os"

	goTerminal "golang.org/x/crypto/ssh/terminal"
)

// Terminal -
type Terminal struct {
	file            *os.File
	theTerminal     *goTerminal.Terminal
	cursortCurrentX int
	cursortCurrentY int
}

// NewTerminal -
func NewTerminal(file *os.File) *Terminal {
	return &Terminal{
		file:        file,
		theTerminal: goTerminal.NewTerminal(file, ""),
	}
}
