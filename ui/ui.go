package ui

import (
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

type Ui interface {
	// print one line to given position.
	PrintLine(str string, x, y int)

	// print lines to given position.
	PrintLines(lines []string, x, y int)
}

type DefaultUi struct {
	// foreground color.
	Fg termbox.Attribute

	// background color.
	Bg termbox.Attribute

	// alert string.
	AlertString string
}

// sugar
func (ui *DefaultUi) Clear() {
	termbox.Clear(ui.Fg, ui.Bg)
}

// sugar
func (ui *DefaultUi) Flush() {
	termbox.Flush()
}

// sugar
func (ui *DefaultUi) Cursor(x int, y int) {
	termbox.SetCursor(x, y)
}

// a function to print N slices to the screen.
func (ui *DefaultUi) PrintLines(lines []string, x, y int) {
	for _, line := range lines {
		ui.PrintLine(line, x, y)
		x = 0
		y++
	}
}

// given a string, this func print them to the terminal,
// return the total width of the string.
func (ui *DefaultUi) PrintLine(str string, x int, y int) int {
	// strWidth is the total width of the string
	strWidth := 0

	width, _ := termbox.Size()

	for len(str) > 0 {
		c, w := utf8.DecodeRuneInString(str)
		if c == utf8.RuneError {
			c = '?'
			w = 1
		}

		str = str[w:]

		termbox.SetCell(x, y, c, ui.Fg, ui.Bg)

		x += runewidth.RuneWidth(c)
		strWidth += runewidth.RuneWidth(c)
	}

	blankPosition := strWidth

	// print blank into the rest of line.
	for blankPosition < width {
		termbox.SetCell(blankPosition, y, ' ', ui.Fg, ui.Bg)
		blankPosition++
	}

	return strWidth
}