package main

import (
	"github.com/weekface/togo/agent"
)

func main() {
	// initialize a new agent object.
	ag := agent.New()

	// run it now.
	ag.Run()
}
