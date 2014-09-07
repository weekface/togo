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
func (agent *Agent) PringLine(str string, x int, y int) int {
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

		termbox.SetCell(x, y, c, agent.Fg, agent.Bg)

		x += runewidth.RuneWidth(c)
		strWidth += runewidth.RuneWidth(c)
	}

	blankPosition := strWidth

	for blankPosition < width {
		termbox.SetCell(blankPosition, y, ' ', agent.Fg, agent.Bg)
		blankPosition++
	}

	return strWidth
}

func (agent *Agent) DeletePromp() {
	_, size := utf8.DecodeLastRuneInString(agent.Chars)
	agent.Chars = agent.Chars[:len(agent.Chars)-size]
	agent.DrawPromp("")
}

// draw the promp line.
func (agent *Agent) DrawPromp(str string) {
	agent.Chars = agent.Chars + string(str)

	prompStr := "TOGO> " + agent.Chars

	x := 0
	y := 0

	width := agent.PringLine(prompStr, x, y)
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
	agent.DrawPromp("")

	agent.PringLine(agent.AlertString, 0, 2)
	agent.PringLine("Latest Version: "+agent.Version, 0, 3)
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
				agent.DrawPromp(str)

			// support back space key
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				agent.DeletePromp()

			// convert rune to string, then draw the promp.
			default:
				str := string(ev.Ch)
				agent.DrawPromp(str)
			}
		}
	}
}
