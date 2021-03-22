package main

//go:generate gotext -srclang=en-US update -out=catalog.go -lang=en,pt,el
import (
	"fmt"
	"github.com/rmateus/uyuni-cli/processor"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"os"
	"os/user"
	"time"
)

// https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/

func main() {
	p := message.NewPrinter(language.Portuguese)
	p.Printf("Hello world!")
	p.Println()
	person := "Alex"
	place := "Utah"
	p.Printf("Hello ", person, " in ", place, "!")
	p.Println()

	p.Printf("sql")
	p.Println()

	for i := 0; i < 2; i++ {
		p.Printf(fmt.Sprint("sql", i))
		p.Println()
	}
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
	p.Println("Goodbye!!")
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
	fmt.Fprintf(logOutput, "[%s] ", time.Now().Format(time.RFC3339))
	defer logOutput.Close()
	for _, arg := range os.Args[1:] {
		fmt.Fprintf(logOutput, "%s ", arg)
	}
	fmt.Fprintln(logOutput, "")
}
