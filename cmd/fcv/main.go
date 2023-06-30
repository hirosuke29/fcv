package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hirosuke29/fcv/internal"
)

// TODO: use cobra CLI framework
func main() {
	flag.Parse()
	subCmd := flag.Arg(0)
	verbose := flag.Bool("debug", false, "enable debug mode")

	switch subCmd {
	case "help":
		help()
	case "init":
		internal.InitRepo(*verbose)
	case "commit":
		if err := internal.Commit(*verbose); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "":
		help()
		os.Exit(1)
	}
}

// help prints help message into stdout.
// TODO: add help message
func help() {
	fmt.Println("Usage: ")
}
