package main

import (
	"github.com/weekface/togo/agent"
	"github.com/weekface/togo/util"
)

func main() {
	// initialize togo data dir.
	util.InitializeTogoDir()

	// initialize a new agent object.
	ag := agent.New()

	// run it now.
	ag.Run()
}
