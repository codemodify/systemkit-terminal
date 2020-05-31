package terminal

// ReadLine -
func (thisRef Terminal) ReadLine() (string, error) {
	return thisRef.theTerminal.ReadLine()
}

// ReadPassword -
func (thisRef Terminal) ReadPassword() (string, error) {
	return thisRef.theTerminal.ReadPassword("")
}
