package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	// [-options]
	helpFlag    bool   // -? / -help
	versionFlag bool   // -version
	cpOption    string // -cp / -classpath

	class string
	args  []string // -Xms<size> ...
}

func parseCmd() *Cmd {

	cmd := &Cmd{}

	flag.Usage = printUsage

	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd

}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
