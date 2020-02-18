package main

import (
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
)

// Exit status
const (
	Good             int = 0
	InvalidOption    int = 1
	InvalidArgument  int = 2
	MissingArguments int = 3
	ExtraArguments   int = 4
	Exception        int = 5
)

// Flags
type cliFlags struct {
	nocolor  bool
	todo     string
	todolist cli.StringList
	verbose  bool
	version  bool
}

var flags cliFlags

const Version = "1.0.0"

func init() {
	// Configure cli package
	cli.Align = true // Defaults to false
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = hl.Sprintf(
		"%s [OPTIONS] <todo1> <todo2>",
		os.Args[0],
	)
	cli.BugEmail = "TODO.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of an error",
			"the exit status will be one of the below:\n\n",
			"1: Invalid option\n",
			"2: Invalid argument\n",
			"3: Missing arguments\n",
			"4: Extra arguments\n",
			"5: Exception",
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
		"Show show stacktrace if error.",
	)
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	hl.Disable(flags.nocolor)

	// Short circuit if version was requested
	if flags.version {
		hl.Printf("TODO version %s\n", Version)
		os.Exit(Good)
	}

	// Validate cli flags
	if cli.NArg() < 2 {
		cli.Usage(MissingArguments)
	} else if cli.NArg() == 2 {
		// TODO
	} else if cli.NArg() > 2 {
		cli.Usage(ExtraArguments)
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			errx(Exception, r.(error).Error())
		}
	}()

	validate()

	// TODO
	for i := range cli.Args() {
		good(cli.Arg(i))
	}

	for i := range flags.todolist {
		warn(flags.todolist[i])
	}
}
