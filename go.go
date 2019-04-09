package main

import (
    "fmt"
    "os"

    "gitlab.com/mjwhitta/cli"
)

var nocolor bool
var todo string
var todolist cli.StringListVar

func init() {
    // Parse cli args
    cli.Banner = fmt.Sprintf("Usage: %s [OPTIONS]", os.Args[0])
    cli.Info = "TODO"
    cli.Bool(
        &nocolor,
        "",
        "nocolor",
        false,
        "Disable colorized outout",
    )
    cli.String(
        &todo,
        "t",
        "todo",
        "TODO",
        "Example for storing cli arg",
    )
    cli.Parse()

    // Validate cli args
    if (cli.NArg() != 0) {
        cli.Usage(1)
    }
}

func main() {
    fmt.Printf("%s\n", todo)
    fmt.Printf("%d %s\n", cli.NArg(), cli.Args())
}
