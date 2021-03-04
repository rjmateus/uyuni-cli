package processor

import (
	"fmt"
	"os"
	"os/exec"
)

type externalToolCmd struct {
	id             string
	description    string
	execPath       string
	providePackage string
}

func (tool externalToolCmd) getId() string {
	return tool.id
}

func (tool externalToolCmd) Execute() error {
	checkCmdError := exec.Command("ls", tool.execPath).Run()
	if checkCmdError != nil {
		fmt.Printf("Unable to run the command '%s'. Check if you have the package '%s' installed\n",
			tool.execPath, tool.providePackage)
		return checkCmdError
	}

	cmd := exec.Command(tool.execPath, os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println(cmd.String())
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (tool externalToolCmd) Info() string {
	return tool.id + ": " + tool.description
}
