package agent

import (
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

// type Agent present the togo terminate's all attributes.
type Agent struct {
	// Now, it contain togo's prompt strings.
	Chars string

	// alert string, it is "Press Ctr+Q to quit.".
	AlertString string

	// foreground color.
	Fg termbox.Attribute

	// background color.
	Bg termbox.Attribute

	// version
	Version string
}

// return a default Agent.
func New() *Agent {
	return &Agent{
		AlertString: "Press what you want. Press Ctr+C to quit.",
		Fg:          termbox.ColorWhite,
		Bg:          termbox.ColorBlack,
		Version:     "0.0.1",
	}
}

// given a string, this func print them to the terminal,
// return the total width of the string.
func (agent *Agent) printLine(str string, x int, y int) int {
	// rw is the total width of the string
	rw := 0

	for len(str) > 0 {
		c, w := utf8.DecodeRuneInString(str)
		if c == utf8.RuneError {
			c = '?'
			w = 1
		}
		str = str[w:]

		termbox.SetCell(x, y, c, agent.Fg, agent.Bg)

		x += runewidth.RuneWidth(c)
		rw += runewidth.RuneWidth(c)
	}
	return rw
}

// draw the promp line.
func (agent *Agent) drawPromp(str string) {
	agent.Chars = agent.Chars + string(str)

	prompStr := "TOGO> " + agent.Chars

	x := 0
	y := 0

	width := agent.printLine(prompStr, x, y)
	termbox.SetCursor(width, 0)
	termbox.Flush()
}

// the public api of the Agent package.
func (agent *Agent) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.Clear(agent.Fg, agent.Bg)
	agent.drawPromp("")

	agent.printLine(agent.AlertString, 0, 2)
	agent.printLine("Latest Version: "+agent.Version, 0, 3)
	termbox.Flush()

loop:

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {

			// quit the program.
			case termbox.KeyCtrlC:
				break loop

			// puts " " to the screen, when we press space.
			case termbox.KeySpace:
				str := " "
				agent.drawPromp(str)

			// convert rune to string, then draw the promp.
			default:
				str := string(ev.Ch)
				agent.drawPromp(str)
			}
		}
	}
}
