package main

func main() {
	cfg := &config{
		commands: getCommands(),
		next:     nil,
		prev:     nil,
	}

	startRepl(cfg)
}
