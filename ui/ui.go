package ui

import (
	"fmt"
	"io"
)

// Ui is an interface for interacting with the terminal,
// Now, there is a DefaultUi implement this interface
type Ui interface {
	// Ask asks the user for input using the given query. The response is
	// returned as the given string, or an error.
	Ask(string) (string, error)

	// Output normal information to console
	Info(string)

	// Output error information to console
	Error(string)
}

// DefaultUi
type DefaultUi struct {
	Reader io.Reader
	Writer io.Writer
}

func (ui *DefaultUi) Ask(question string) (string, error) {
	fmt.Println("Asking " + question)
	return "", nil
}

func (ui *DefaultUi) Info(message string) {
	fmt.Fprint(ui.Writer, message)
	fmt.Fprint(ui.Writer, "\n")
}

func (ui *DefaultUi) Error(message string) {
}
