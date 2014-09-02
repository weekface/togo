package agent

import (
	"fmt"
	"os"

	"github.com/weekface/togo/ui"
)

type Agent struct {
	Ui   ui.DefaultUi
	Args map[string]interface{}
}

func New(args map[string]interface{}) *Agent {
	ui := ui.DefaultUi{Reader: os.Stdin, Writer: os.Stdout}
	return &Agent{Ui: ui, Args: args}
}

func (agent *Agent) Run() {
	if agent.Args["list"].(bool) {
		fmt.Println("listing...")
	}
}
