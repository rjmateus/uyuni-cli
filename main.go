package main

import (
	"fmt"
	"os"

	"github.com/rmateus/uyuni-cli/processor"
)

// https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/

func main() {
	fmt.Println("Hello, I'm here to serve you!!")
	manager := processor.GetToolsCommandManager()
	if len(os.Args) < 2 {
		manager.UsagePrint()
		os.Exit(1)
	}

	manager.Execute(os.Args[1])

	fmt.Println("Goodbye!!")
}
