package main

import (
	"fmt"
	"github.com/rmateus/sumatools/newTool"
	"os"
	"os/exec"
)
// https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/

const (
	usage = `SUMA tools 
Usage should be "sumatools [command]"

Available Command:
`
)
var proxyCommands = map[string]string{
"spacewalk-sql": "/usr/bin/spacewalk-sql",
"spacewalk-repo-sync": "/usr/bin/spacewalk-repo-sync",
"satpasswd": "/usr/bin/satpasswd",
}

var commands = []Command{
	{id:"spacewalk-sql",
		localProcess: func() {runCommand("/usr/bin/spacewalk-sql")	}},
	{id:"spacewalk-sql",localProcess: newTool.ProcessNewTool},
}

func usagePrint(){
	fmt.Println(usage)
	for k, _ := range proxyCommands {
		fmt.Println("  - ", k)
	}
}

func main() {
	fmt.Println("Hello, I'm here to serve you!!")
	if len(os.Args) < 2 {
		usagePrint()
		os.Exit(1)
	}
	value, ok:= proxyCommands[os.Args[1]]
	if ok {
		runCommand(value)
	}else if os.Args[1] == "newTool" {
		newTool.ProcessNewTool()
	}else{
		usagePrint()
		os.Exit(1)
	}

	fmt.Println("My work in here is done!!")
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