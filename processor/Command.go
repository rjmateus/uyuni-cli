package processor

import (
	"fmt"
	"os"
	"os/exec"
)

type toolCmd struct{
	Id  string
	description string
	run func()
}

func (tool toolCmd) Execute() {
	tool.run()
}

// internationalization
// https://pkg.go.dev/github.com/cloudfoundry-attic/jibber_jabber
// https://phrase.com/blog/posts/internationalization-i18n-go/

// alternative: https://github.com/kataras/i18n
func (tool *toolCmd) Info() string{
	return tool.Id + ": " + tool.description
}


func externalToolCommand(id string, execPath string, desc string) toolCmd {
	return toolCmd{Id:id, description: desc,
		run: func() {runCommand(execPath)	}}
}

func internalToolCommand(id string, ft func(), desc string) toolCmd {
	return toolCmd{Id:id, description: desc, run: ft}
}

func runCommand(command string)  {
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