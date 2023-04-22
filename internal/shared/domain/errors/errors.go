package errors

// CommandNotRegisteredError will return an error when an command not registered
type CommandNotRegisteredError string

// CommandNotRegisteredError implements the Error interface
func (e CommandNotRegisteredError) Error() string {
	return string(e)
}

// CommandAlreadyRegisteredError will return an error when an command not registered
type CommandAlreadyRegisteredError string

// CommandAlreadyRegisteredError implements the Error interface
func (e CommandAlreadyRegisteredError) Error() string {
	return string(e)
}
