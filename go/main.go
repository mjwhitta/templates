package main

import (
	"gitlab.com/mjwhitta/cli"
	"gitlab.com/mjwhitta/log"
)

// Exit status
const (
	Good = iota
	InvalidOption
	InvalidArgument
	MissingArguments
	ExtraArguments
	Exception
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			log.ErrX(Exception, r.(error).Error())
		}
	}()

	validate()

	// TODO
	for i := range cli.Args() {
		log.Good(cli.Arg(i))
	}

	for i := range flags.todolist {
		log.Warn(flags.todolist[i])
	}
}
