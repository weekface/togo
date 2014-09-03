package main

import (
	"github.com/docopt/docopt-go"
	"github.com/weekface/togo/agent"
)

const Usage = `togo cli.

Usage:
  togo list
  togo add    <message>
  togo rm     <id>
  togo finish <id>
  togo show   <id>
  togo server <port> [-d|--deamon]
  togo -h | --help
  togo --version

Options:
  -d --deamon   Start the web server as a deamon.
  -h --help     Show this screen.
  --version     Show version.`

func main() {
	arguments, err := docopt.Parse(Usage, nil, true, "Togo 0.0.1", false)
	if err != nil {
		panic(err)
	}

	ag := agent.New(arguments)
	ag.Run()
}
