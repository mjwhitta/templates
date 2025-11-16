package main

import (
	"github.com/mjwhitta/cli"
	"github.com/mjwhitta/log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r)
			}

			switch r := r.(type) {
			case error:
				log.ErrX(Exception, r.Error())
			case string:
				log.ErrX(Exception, r)
			}
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
