package newTool

import (
	"flag"
	"fmt"
	"os"
)
const (
	usage = `My new tool

Usage: myTool [Options]
	

Options:
`
)

func ProcessNewTool(){
	fmt.Println("my supper new local tool")
	flag.Usage = func() {
		fmt.Println(usage)
		flag.PrintDefaults()
	}
	myString := flag.String("path", ".", "Location for generated data")

	myBool := flag.Bool("b", false, "Create dot file for Graphviz")

	if len(os.Args) < 3 {
		flag.Usage()
		os.Exit(1)
	}

	flag.CommandLine.Parse(os.Args[2:])
	fmt.Println("paramters:")
	fmt.Printf("myString: %s\n", *myString)
	fmt.Printf("myBool: %t\n", *myBool)
}