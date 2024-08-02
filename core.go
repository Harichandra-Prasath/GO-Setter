package main

import (
	"fmt"
	"os"
)

type Config struct {
	Name     string
	Root     string
	Origin   string
	Branch   string
	IsIgnore bool
	IsMake   bool
	IsReadme bool
	Mod      string
}

func NewCfg() *Config {

	// Set Default Values
	var cfg Config

	home, err := os.UserHomeDir()
	chkerr(err)

	cfg.Root = home
	cfg.Branch = "main" // default git branch
	cfg.IsIgnore = true
	cfg.IsMake = true
	cfg.IsReadme = true

	return &cfg
}

func Prompt(cfg *Config) {
	fmt.Println("\nWelcome to GO-Setter!!!")
	fmt.Println()

	var name string
	fmt.Print("Enter your Project Name: ")
	n, _ := fmt.Scanln(&name)
	if n == 0 {
		panic("Project name is required")
	}
	if len(name) < 3 {
		panic("Project Name should be 3 Characters atleast")
	}
	cfg.Name = name

	var root string
	fmt.Print("Enter your Root Directory (default is home): ")
	n, _ = fmt.Scanln(&root)
	if n != 0 {
		_, err := os.ReadDir(root)
		if os.IsNotExist(err) {
			panic("Given Directory doesnt exist and cannot be used as root")
		}
		cfg.Root = root
	}

	var origin string
	fmt.Print("Enter the remote origin: ")
	n, _ = fmt.Scanln(&origin)
	if n == 0 {
		panic("Remote repo link is required")
	}
	cfg.Origin = origin

	var branch string
	fmt.Print("Enter the branch to move/rename from master (default is main): ")
	n, _ = fmt.Scanln(&branch)
	if n != 0 {
		cfg.Branch = branch
	}

	var mod string
	fmt.Print("Enter go module path (default will be extracted from your remote origin): ")
	n, _ = fmt.Scanln(&mod)
	if n != 0 {
		cfg.Mod = mod
	} else {
		mod, err := getMod(cfg.Origin)
		chkerr(err)
		cfg.Mod = mod
	}

	var ignore string
	fmt.Print("Do you need a minimal .gitignore? (yes or no) (default is yes): ")
	n, _ = fmt.Scanln(&ignore)
	if n != 0 {
		updateBool(ignore, &cfg.IsIgnore)
	}

	var make string
	fmt.Print("Do you need a minimal Makefile? (yes or no) (default is yes): ")
	n, _ = fmt.Scanln(&make)
	if n != 0 {
		updateBool(make, &cfg.IsMake)
	}

	var readme string
	fmt.Print("Do you need a readme? (yes or no) (default is yes): ")
	n, _ = fmt.Scanln(&readme)
	if n != 0 {
		updateBool(readme, &cfg.IsMake)
	}

}
