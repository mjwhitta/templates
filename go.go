package main

import (
    "fmt"
    "os"

    "gitlab.com/mjwhitta/cli"
)

// Helpers begin
func err(msg string) {
    var color = "\x1b[31m"
    if (nocolor) {color = ""}
    fmt.Printf("%s[!] %s\x1b[0m\n", color, msg)
}
func errx(status int, msg string) {
    err(msg)
    os.Exit(status)
}
func good(msg string) {
    var color = "\x1b[32m"
    if (nocolor) {color = ""}
    fmt.Printf("%s[+] %s\x1b[0m\n", color, msg)
}
func info(msg string) {
    var color = "\x1b[37m"
    if (nocolor) {color = ""}
    fmt.Printf("%s[*] %s\x1b[0m\n", color, msg)
}
func subinfo(msg string) {
    var color = "\x1b[36m"
    if (nocolor) {color = ""}
    fmt.Printf("%s[=] %s\x1b[0m\n", color, msg)
}
func warn(msg string) {
    var color = "\x1b[33m"
    if (nocolor) {color = ""}
    fmt.Printf("%s[-] %s\x1b[0m\n", color, msg)
}
// Helpers end

var nocolor bool
var todo string
var todolist cli.StringListVar

func init() {
    // Parse cli args
    cli.Banner = fmt.Sprintf(
        "Usage: %s [OPTIONS] <arg1>... [argN]",
        os.Args[0],
    )
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
    if (cli.NArg() == 0) {
        cli.Usage(1)
    }
}

func main() {
    for i := range cli.Args() {
        good(cli.Arg(i))
    }
}
