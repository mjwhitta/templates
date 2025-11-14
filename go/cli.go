package main

import (
	"os"
	"path/filepath"

	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
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
	cli.Banner = "" +
		filepath.Base(os.Args[0]) + " [OPTIONS] <todo1> <todo2>"
	cli.BugEmail = "TODO.bugs@whitta.dev"

	cli.ExitStatus(
		"Normally the exit status is 0. In the event of an error the",
		"exit status will be one of the below:\n\n",
		hl.Sprintf("  %d: Invalid option\n", InvalidOption),
		hl.Sprintf("  %d: Missing option\n", MissingOption),
		hl.Sprintf("  %d: Invalid argument\n", InvalidArgument),
		hl.Sprintf("  %d: Missing argument\n", MissingArgument),
		hl.Sprintf("  %d: Extra argument\n", ExtraArgument),
		hl.Sprintf("  %d: Exception", Exception),
	)
	cli.Info("TODO")

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
	hl.Disable(flags.nocolor)

	// Short circuit, if version was requested
	if flags.version {
		hl.Printf(
			"%s version %s\n",
			filepath.Base(os.Args[0]),
			Version,
		)
		os.Exit(Good)
	}

	// Validate cli flags
	switch {
	case cli.NArg() < 2: //nolint:mnd // 2 cli args
		cli.Usage(MissingArgument)
	case cli.NArg() == 2: //nolint:mnd // 2 cli args
	// TODO
	case cli.NArg() > 2: //nolint:mnd // 2 cli args
		cli.Usage(ExtraArgument)
	}
}
