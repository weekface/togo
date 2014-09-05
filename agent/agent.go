package agent

import (
	"os"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
	"github.com/weekface/togo/ui"
)

func setString(str string, fg, bg termbox.Attribute, x int, y int) int {
	rw := 0
	for len(str) > 0 {
		c, w := utf8.DecodeRuneInString(str)
		if c == utf8.RuneError {
			c = '?'
			w = 1
		}
		str = str[w:]
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
		rw += runewidth.RuneWidth(c)
	}
	return rw
}

type Agent struct {
	Ui    ui.DefaultUi
	Chars string
}

func New() *Agent {
	ui := ui.DefaultUi{Reader: os.Stdin, Writer: os.Stdout}
	return &Agent{Ui: ui}
}

func (agent *Agent) drawPromp(str string) {
	agent.Chars = agent.Chars + string(str)

	prompStr := "TOGO> " + agent.Chars

	bg := termbox.ColorBlack
	fg := termbox.ColorWhite
	x := 0
	y := 0

	termbox.Clear(fg, bg)
	width := setString(prompStr, fg, bg, x, y)
	termbox.SetCursor(width, 0)
	setString("Press Ctr+Q to quit.", termbox.ColorBlack, termbox.ColorWhite, 0, 2)
	termbox.Flush()
}

func (agent *Agent) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	agent.drawPromp("")

	setString("Press Ctr+Q to quit.", termbox.ColorBlack, termbox.ColorWhite, 0, 2)
	termbox.Flush()

loop:

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlQ:
				break loop
			default:
				str := string(ev.Ch)
				agent.drawPromp(str)
			}
		}
	}
}
