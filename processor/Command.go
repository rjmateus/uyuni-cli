package processor

import (
	"fmt"
	"os"
	"os/exec"
)

type toolCmd struct {
	Id          string
	description string
	run         func()
}

func (tool toolCmd) Execute() {
	tool.run()
}

func (tool *toolCmd) Info() string {
	return tool.Id + ": " + tool.description
}

func externalToolCommand(id string, execPath string, desc string) toolCmd {
	return toolCmd{Id: id, description: desc,
		run: func() { runCommand(execPath) }}
}

func runCommand(command string) {
	cmd := exec.Command(command, os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println(cmd.String())
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
