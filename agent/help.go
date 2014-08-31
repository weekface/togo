package agent

func (agent *Agent) Help() {
	helpText := `
Usage: togo [action] [options]

  Start use togo command. But now, this is just a demooooooo...

Actions:
	
	add     Add a todo
	finsh   Finish a todo
	delete  Delete a todo
	edit    Update a todo

Options:
	
	--server          Start a httpserver also
	--port            The http server port
	--date-dir= path  Path to a data directory to store todo data
	=-pid-file=path   Path to file to store agent PID

Press Ctr+C to quit.
`
	agent.Ui.Info(helpText)
}
