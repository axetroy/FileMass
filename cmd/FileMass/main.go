package main

import (
	"flag"
	"fmt"
	"os"

	fileMass "github.com/axetroy/FileMass"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func printHelp() {
	println(`FileMass - a powerful file generation tool used to create directory structures and various types of files

USAGE:
  filemass [OPTIONS]

OPTIONS:
  --help                Print help information
  --version             Print version information
	--concurrence         Set the number of concurrent tasks
	--output              Set the output directory
	--depth								Set the depth of the directory
	--min-size						Set the minimum size of the file
	--max-size						Set the maximum size of the file
	--dirs								Set the number of directories
	--files								Set the number of files in the directory

SOURCE CODE:
  https://github.com/axetroy/FileMass`)
}

func run() error {
	var (
		showHelp    bool
		showVersion bool
		noColor     bool
	)

	flag.BoolVar(&noColor, "no-color", false, "disabled color for printing")
	flag.BoolVar(&showHelp, "help", false, "Print help information")
	flag.BoolVar(&showVersion, "version", false, "Print version information")

	flag.Usage = printHelp

	flag.Parse()

	if showHelp {
		printHelp()
		os.Exit(0)
	}

	if showVersion {
		println(fmt.Sprintf("%s %s %s", version, commit, date))
		os.Exit(0)
	}

	fileMass.Mass()

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(255)
	}
}
