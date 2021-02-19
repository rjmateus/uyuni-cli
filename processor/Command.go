package processor

import (
	"fmt"
	"os"
	"os/exec"
)

type toolCmd struct{
	Id  string
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
	return tool.Id
}


func remoteToolCommand(id string, execPath string) toolCmd {
	return toolCmd{Id:id,
		run: func() {runCommand(execPath)	}}
}

func localToolCommand(id string, ft func()) toolCmd {
	return toolCmd{Id:id, run: ft}
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