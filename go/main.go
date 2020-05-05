package main

import (
	"gitlab.com/mjwhitta/cli"
	"gitlab.com/mjwhitta/log"
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
