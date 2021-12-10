package main

import (
	"fmt"
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
)

// Exit status
const (
	Good = iota
	InvalidOption
	MissingOption
	InvalidArgument
	MissingArgument
	ExtraArgument
	Exception
)

// Flags
var flags struct {
	nocolor  bool
	todo     string
	todolist cli.StringList
	verbose  bool
	version  bool
}

func init() {
	// Configure cli package
	cli.Align = true // Defaults to false
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = fmt.Sprintf(
		"%s [OPTIONS] <todo1> <todo2>",
		os.Args[0],
	)
	cli.BugEmail = "TODO.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of an error",
			"the exit status will be one of the below:\n\n",
			fmt.Sprintf("%d: Invalid option\n", InvalidOption),
			fmt.Sprintf("%d: Missing option\n", MissingOption),
			fmt.Sprintf("%d: Invalid argument\n", InvalidArgument),
			fmt.Sprintf("%d: Missing argument\n", MissingArguments),
			fmt.Sprintf("%d: Extra argument\n", ExtraArguments),
			fmt.Sprintf("%d: Exception", Exception),
		},
		" ",
	)
	cli.Info = strings.Join([]string{"TODO."}, " ")
	// cli.MaxWidth = 80 // Defaults to 80
	cli.SeeAlso = []string{"TODO"}
	// cli.TabWidth = 4 // Defaults to 4
	cli.Title = "TODO"

	// Parse cli flags
	cli.Flag(
		&flags.nocolor,
		"no-color",
		false,
		"Disable colorized output.",
	)
	cli.Flag(&flags.todo, "t", "todo", "TODO", "Describe TODO.")
	cli.Flag(&flags.todolist, "todolist", "Describe TODOlist.")
	cli.Flag(
		&flags.verbose,
		"v",
		"verbose",
		false,
		"Show stacktrace, if error.",
	)
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	fmt.Disable(flags.nocolor)

	// Short circuit if version was requested
	if flags.version {
		fmt.Printf("TODO version %s\n", Version)
		os.Exit(Good)
	}

	// Validate cli flags
	if cli.NArg() < 2 {
		cli.Usage(MissingArguments)
	} else if cli.NArg() == 2 {
		// TODO
	} else if cli.NArg() > 2 {
		cli.Usage(ExtraArgument)
	}
}
