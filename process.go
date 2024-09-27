package main

import (
	"os"
	"os/exec"
)

func Process(cfg *Config) {

	// Stage 1 (DIR Stage)

	chkerr(os.Chdir(cfg.Root))

	chkerr(os.Mkdir(cfg.Name, 0757))

	chkerr(os.Chdir(cfg.Name))

	// Stage 2 (GIT Stage)

	cmd := exec.Command("git", "init")
	chkerr(cmd.Run())

	cmd = exec.Command("git", "branch", "-m", cfg.Branch)
	chkerr(cmd.Run())

	cmd = exec.Command("git", "remote", "add", "origin", cfg.Origin)
	chkerr(cmd.Run())

	// Stage 3 (GO Stage)

	cmd = exec.Command("go", "mod", "init", cfg.Mod)
	chkerr(cmd.Run())

	// Stage 4 (File Stage)

	if cfg.IsReadme {
		chkerr(os.WriteFile("readme.md", getRContent(cfg.Name), 0757))
	}
	if cfg.IsMake {
		chkerr(os.WriteFile("Makefile", getMContent(cfg.Name), 0757))
	}
	if cfg.IsIgnore {
		chkerr(os.WriteFile(".gitignore", getIcontent(), 0757))
	}

	chkerr(os.WriteFile("main.go", getMaincontent(), 0757))

}
