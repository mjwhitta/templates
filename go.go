package main

import (
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
)

const Version = "1.0.0"

// Helpers begin

func err(msg string) { hl.PrintlnRed("[!] %s", msg) }
func errx(status int, msg string) {
	err(msg)
	os.Exit(status)
}
func good(msg string)    { hl.PrintlnGreen("[+] %s", msg) }
func info(msg string)    { hl.PrintlnWhite("[*] %s", msg) }
func subinfo(msg string) { hl.PrintlnCyan("[=] %s", msg) }
func warn(msg string)    { hl.PrintlnYellow("[-] %s", msg) }

// Helpers end

var nocolor bool
var todo string
var todolist cli.StringList
var version bool

func init() {
	// Configure cli package
	cli.Align = false // Defaults to false
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = hl.Sprintf(
		"%s [OPTIONS] <arg1>... [argN]",
		os.Args[0],
	)
	cli.BugEmail = "todo.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of invalid",
			"or missing arguments, the exit status will be non-zero.",
		},
		" ",
	)
	cli.Info = strings.Join([]string{"TODO"}, " ")
	cli.MaxWidth = 80 // Defaults to 80
	// cli.SeeAlso = []string{"TODO"}
	cli.TabWidth = 4 // Defaults to 4
	cli.Title = "TODO"

	// Parse cli flags
	cli.Flag(&nocolor, "no-color", false, "Disable colorized output.")
	cli.Flag(&todo, "t", "todo", "TODO", "Describe TODO.")
	cli.Flag(&todolist, "todolist", "Describe TODOlist.")
	cli.Flag(&version, "V", "version", false, "Show version.")
	cli.Parse()

	// Validate cli args (todo)
	if !version && (cli.NArg() == 0) {
		cli.Usage(1)
	}
}

func main() {
	hl.Disable = nocolor

	defer func() {
		if r := recover(); r != nil {
			err(r.(error).Error())
		}
	}()

	if version {
		hl.Println("Version: " + Version)
	} else {
		// TODO
		for i := range cli.Args() {
			good(cli.Arg(i))
		}

		for i := range todolist {
			warn(todolist[i])
		}
	}
}
