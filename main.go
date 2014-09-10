package main

import (
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/weekface/togo/agent"
)

func main() {
	dir, _ := homedir.Dir()
	togoDir := filepath.Join(dir, ".togo")
	os.MkdirAll(filepath.Join(togoDir, "new"), 0755)
	os.MkdirAll(filepath.Join(togoDir, "old"), 0755)

	// initialize a new agent object.
	ag := agent.New()

	// run it now.
	ag.Run()
}
