package main

import (
    "fmt"
    "os"

    "gitlab.com/mjwhitta/cli"
    hl "gitlab.com/mjwhitta/hilighter/src"
)

const Version = "1.0.0"

// Helpers begin
func err(msg string) {fmt.Print(hl.Red("[!] %s\n", msg))}
func errx(status int, msg string) {err(msg); os.Exit(status)}
func good(msg string) {fmt.Print(hl.Green("[+] %s\n", msg))}
func info(msg string) {fmt.Print(hl.White("[*] %s\n", msg))}
func subinfo(msg string) {fmt.Print(hl.Cyan("[=] %s\n", msg))}
func warn(msg string) {fmt.Print(hl.Yellow("[-] %s\n", msg))}
// Helpers end

var nocolor bool
var todo string
var todolist cli.StringList
var version bool

func init() {
    // Configure cli package
    cli.Align = true
    cli.Banner = fmt.Sprintf(
        "Usage: %s [OPTIONS] <arg1>... [argN]",
        os.Args[0],
    )
    cli.Info = "TODO"

    // Parse cli args
    cli.Flag(&nocolor, "no-color", false, "Disable colorized outout")
    cli.Flag(&todo, "t", "todo", "TODO", "Describe TODO")
    cli.Flag(&version, "V", "version", false, "Show version")
    cli.Parse()

    // Validate cli args
    if (!version && (cli.NArg() == 0)) {
        cli.Usage(1)
    }
}

func main() {
    if (version) {
        fmt.Printf("Version: %s\n", Version)
    } else {
        // TODO
        for i := range cli.Args() {
            good(cli.Arg(i))
        }
    }
}
