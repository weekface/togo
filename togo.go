package main

import (
	"os"

	"github.com/weekface/togo/agent"
)

func main() {
	ag := agent.New(os.Args[1:])
	ag.Run()
}
