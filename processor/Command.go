package processor

import (
	"fmt"
	"os"
	"os/exec"
)

type ToolCmd struct{
	Id  string
	run func()
}

func (tool *ToolCmd) Execute() {
	tool.run()
}

// internationalization
// https://pkg.go.dev/github.com/cloudfoundry-attic/jibber_jabber
// https://phrase.com/blog/posts/internationalization-i18n-go/

// alternative: https://github.com/kataras/i18n
func (tool *ToolCmd) Info() string{
	return tool.Id
}


func RemoteToolCommand(id string, execPath string) ToolCmd {
	return ToolCmd{Id:id,
		run: func() {runCommand(execPath)	}}
}

func LocalToolCommand(id string, ft func()) ToolCmd {
	return ToolCmd{Id:id, run: ft}
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