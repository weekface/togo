package agent

import (
	"os"

	"github.com/nsf/termbox-go"
	"github.com/weekface/togo/ui"
)

var help = `
Usage:
	
  togo --server

Options:
	
  -s --server  Start server also.
	
Press Ctr+Q to quit.
`

func setString(str string, fg, bg termbox.Attribute) {
	cols, lines := termbox.Size()

	x := 0
	y := 0

	for _, r := range str {
		if x >= cols {
			x = 0
			y++
		}
		if y >= lines {
			break
		}
		if r != '\n' {
			termbox.SetCell(x, y, r, fg, bg)
			x++
		} else {
			y++
			x = 0
		}
	}
}

func drawHelp() {
	bg := termbox.ColorBlack
	fg := termbox.ColorWhite

	termbox.Clear(fg, bg)
	setString(help, fg, bg)
	termbox.Flush()
}

type Agent struct {
	Ui ui.DefaultUi
}

func New() *Agent {
	ui := ui.DefaultUi{Reader: os.Stdin, Writer: os.Stdout}
	return &Agent{Ui: ui}
}

func (agent *Agent) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	drawHelp()

loop:

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyCtrlQ {
				break loop
			}
		}
	}
}
