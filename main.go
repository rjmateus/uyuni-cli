package main

import (
	"fmt"
	"github.com/rmateus/uyuni-cli/processor"
	"log"
	"os"
	"os/user"
	"time"
)

// https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/

func main() {
	manager := processor.GetToolsCommandManager()
	if len(os.Args) < 2 {
		manager.UsagePrint()
		os.Exit(1)
	}

	logCommandHistory()
	err := manager.Execute(os.Args[1])
	if err != nil {
		os.Exit(1)
	}
}

func logCommandHistory() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	logOutput, err := os.OpenFile(usr.HomeDir+"/.uyuni_cli_history", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(logOutput, "[%s]", time.Now().Format(time.RFC3339))
	defer logOutput.Close()
	time.Now()
	for _, arg := range os.Args[1:] {
		fmt.Fprintf(logOutput, arg+" ")
		fmt.Print(arg + " ")
	}
	fmt.Fprintln(logOutput, "")
}
