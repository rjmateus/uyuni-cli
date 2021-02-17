package main

import (
	"fmt"
	"github.com/rmateus/sumatools/newTool"
	"github.com/rmateus/sumatools/processor"
	"os"
)
// https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/

const (
	usage = `SUMA tools 
Usage should be "sumatools [command]"

Available Command:`
)

var commands = map[string]processor.ToolCmd{
	"spacewalk-sql": processor.RemoteToolCommand("spacewalk-sql", "/usr/bin/spacewalk-sql"),
	"spacewalk-repo-sync": processor.RemoteToolCommand("spacewalk-repo-sync", "/usr/bin/spacewalk-repo-sync"),
	"satpasswd": processor.RemoteToolCommand("satpasswd", "/usr/bin/satpasswd"),
	"newTool": processor.LocalToolCommand("newTool", newTool.ProcessNewTool),
}

func usagePrint(){
	fmt.Println(usage)
	for _, value := range commands {
		fmt.Println("  - ", value.Info())
	}
}

func main() {
	fmt.Println("Hello, I'm here to serve you!!")
	if len(os.Args) < 2 {
		usagePrint()
		os.Exit(1)
	}

	if value, ok:= commands[os.Args[1]]; ok {
		value.Execute()
	}else{
		usagePrint()
		os.Exit(1)
	}

	fmt.Println("My work in here is done!!")
}
