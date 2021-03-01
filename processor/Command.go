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

func (tool *toolCmd) Execute() {
	tool.run()
}

func (tool *toolCmd) Info() string {
	return tool.Id + ": " + tool.description
}

func externalToolCommand(id string, execPath string, desc string, providePackage string) toolCmd {
	return toolCmd{Id: id, description: desc,
		run: func() { runCommand(execPath, providePackage) }}
}

func runCommand(command string, providePackage string) {
	checkCmdError := exec.Command("ls", command).Run()
	if checkCmdError != nil {
		fmt.Printf("Unable to run the command '%s'. Check if you have the package '%s' installed\n", command, providePackage)
		return
	}

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
