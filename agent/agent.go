package agent

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/mitchellh/go-homedir"
	"github.com/nsf/termbox-go"
	"github.com/weekface/togo/ui"
	"github.com/weekface/togo/util"
)

var help = `
Usage:

  list|l        List all the todo s.
  add <task>    Add a new todo task.
  finish <id>   Finish a todo task.
  help|h        Show help info.
  quit|exit|q   Quit TOGO.

Press Ctr+C to quit.`

// type Agent present the togo terminate's all attributes.
type Agent struct {
	Ui *ui.DefaultUi

	// Now, it contain togo's prompt strings.
	Chars string

	// version
	Version string

	// new files path
	NewPath string

	// old files path
	OldPath string

	// all new todos.
	lists []ui.Todo
}

// return a default Agent.
func NewAgent() *Agent {
	dir, _ := homedir.Dir()
	return &Agent{
		Ui: &ui.DefaultUi{
			Fg:          termbox.ColorDefault,
			Bg:          termbox.ColorDefault,
			AlertString: "Press what you want. Press Ctr+C to quit.",
		},
		Version: "1.0.0",
		NewPath: filepath.Join(dir, ".togo/new"),
		OldPath: filepath.Join(dir, ".togo/old"),
	}
}

// backspace key support
func (agent *Agent) DeletePromp() {
	_, size := utf8.DecodeLastRuneInString(agent.Chars)
	agent.Chars = agent.Chars[:len(agent.Chars)-size]
	agent.DrawPromp("")
}

// sugar
func (agent *Agent) Clear() {
	agent.Ui.Clear()
}

// sugar
func (agent *Agent) Flush() {
	agent.Ui.Flush()
}

// display help info
func (agent *Agent) ShowHelp() {
	agent.Clear()

	agent.Chars = ""
	agent.DrawPromp("")

	lines := strings.Split(help, "\n")

	x := 1
	y := 1

	for _, line := range lines {
		agent.Ui.PrintLine(line, x, y)
		x = 0
		y++
	}

	agent.Flush()
}

// find all undo todos
func (agent *Agent) FindNewTodos() {
	files, err := filepath.Glob(filepath.Join(agent.NewPath, "*"))

	// youxi: http://golang.org/src/pkg/sort/sort.go?s=6285:6310#L248
	sort.Sort(sort.Reverse(sort.StringSlice(files)))

	list := make([]ui.Todo, 0)

	if err != nil {
		panic("Read todos fail!")
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		list = append(list, ui.NewTodo(string(data), file))
	}

	agent.lists = list
}

// display them.
func (agent *Agent) DisplayAll() {
	agent.Clear()
	agent.Ui.PrintLines(agent.lists, 0, 2)
	agent.Chars = ""
	agent.DrawPromp("")
	agent.Flush()
}

// ok, this is owaful.
func (agent *Agent) FindAndDisplayAll() {
	agent.FindNewTodos()
	agent.DisplayAll()
}

// add a todo.
func (agent *Agent) Add() {
	slice := strings.SplitN(agent.Chars, " ", 2)

	if len(slice) == 2 {
		filename := util.Hash()
		ioutil.WriteFile(filepath.Join(agent.NewPath, filename), []byte(slice[1]), 0644)
		agent.FindAndDisplayAll()
	}
}

// finish a todo.
func (agent *Agent) finish() {
	slice := strings.SplitN(agent.Chars, " ", 2)

	if len(slice) != 2 {
		return
	}

	index, err := strconv.Atoi(slice[1])

	if err != nil {
		return
	}

	todo := agent.lists[index-1]
	err = os.Remove(todo.Filename)

	if err == nil {
		agent.FindAndDisplayAll()
	}
}

// quit.
func (agent *Agent) Quit() {
	termbox.Close()
	os.Exit(0)
}

// parse command. Now, just support help command to show the help info.
func (agent *Agent) ParseCmd() {
	command := strings.Split(agent.Chars, " ")[0]

	switch command {

	// help command.
	case "help", "h":
		agent.ShowHelp()

	// quit command.
	case "quit", "Quit", "Exit", "exit", "q":
		agent.Quit()

	// list command.
	case "list", "l":
		agent.FindAndDisplayAll()

	// add command.
	case "add", "Add":
		agent.Add()

	// finish command.
	case "finish":
		agent.finish()

	// default, print what you press.
	default:
		agent.ShowHelp()
	}
}

// draw the promp line.
func (agent *Agent) DrawPromp(str string) {
	agent.Chars = agent.Chars + string(str)

	prompStr := "TOGO> " + agent.Chars

	x := 0
	y := 0

	width := agent.Ui.PrintLine(prompStr, x, y)
	agent.Ui.Cursor(width, 0)
	agent.Flush()
}

// the public api of the Agent package.
func (agent *Agent) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	agent.Clear()
	agent.DrawPromp("")

	agent.FindAndDisplayAll()
	agent.Flush()

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

			// when we press Enter, parse the command.
			case termbox.KeyEnter:
				agent.ParseCmd()

			// convert rune to string, then draw the promp.
			default:
				str := string(ev.Ch)
				agent.DrawPromp(str)
			}
		}
	}
}
