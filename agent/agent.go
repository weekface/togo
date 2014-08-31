package agent

import (
	"os"

	gracefully "github.com/visionmedia/go-gracefully"
	"github.com/weekface/togo/ui"
)

type Agent struct {
	Ui ui.DefaultUi
}

func New(args []string) *Agent {
	ui := ui.DefaultUi{Reader: os.Stdin, Writer: os.Stdout}
	return &Agent{Ui: ui}
}

func (agent *Agent) Run() {
	agent.Help()
	gracefully.Shutdown()
}
