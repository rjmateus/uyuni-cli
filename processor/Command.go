package processor

import (
	"fmt"
	"os"
	"os/exec"
)

//type Tools struct {
//	Tools map[string]ToolCmd
//}
//
//func (t *Tools) RegisterTool(tool ToolCmd)  {
//	t.Tools[tool.Id] = tool
//}

type ToolCmd struct{
	Id  string
	run func()
}

func (tool *ToolCmd) Execute() {
	tool.run()
}

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