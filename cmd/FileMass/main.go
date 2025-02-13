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
  --concurrence         Set the number of concurrent tasks, default is 1
  --output              Set the output directory
  --depth               Set the depth of the directory, default is 1
  --min-size            Set the minimum size(in KB) of the file, default is 1KB
  --max-size            Set the maximum size(in KB) of the file, default is 1024KB
  --dirs                Set the number of directories, default is 1
  --files               Set the number of files in the directory, default is 1
  --clean               Clean the output directory before generating files

SOURCE CODE:
  https://github.com/axetroy/FileMass`)
}

func run() error {
	var (
		showHelp    bool
		showVersion bool
		concurrence int
		output      string
		depth       int
		minSize     int
		maxSize     int
		dirs        int
		files       int
		clean       bool
	)

	flag.BoolVar(&showHelp, "help", false, "Print help information")
	flag.BoolVar(&showVersion, "version", false, "Print version information")
	flag.IntVar(&concurrence, "concurrence", 1, "Set the number of concurrent tasks")
	flag.StringVar(&output, "output", "", "Set the output directory")
	flag.IntVar(&depth, "depth", 1, "Set the depth of the directory")
	flag.IntVar(&minSize, "min-size", 1, "Set the minimum size of the file (in KB)")
	flag.IntVar(&maxSize, "max-size", 1024, "Set the maximum size of the file (in KB)")
	flag.IntVar(&dirs, "dirs", 1, "Set the number of directories")
	flag.IntVar(&files, "files", 1, "Set the number of files in the directory")
	flag.BoolVar(&clean, "clean", false, "Clean the output directory before generating files")

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

	if output == "" {
		return fmt.Errorf("--output is required")
	}

	config := fileMass.Config{
		Concurrence: concurrence,
		Output:      output,
		Depth:       depth,
		MinSize:     minSize,
		MaxSize:     maxSize,
		Dirs:        dirs,
		Files:       files,
		Clean:       clean,
	}

	return fileMass.Mass(config)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(255)
	}
}
