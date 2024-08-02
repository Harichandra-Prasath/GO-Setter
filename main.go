package main

func main() {

	// get a new cfg with default values
	cfg := NewCfg()

	Prompt(cfg)
	Process(cfg)
}
