package main

import (
	"github.com/mjwhitta/cli"
	"github.com/mjwhitta/log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch r := r.(type) {
			case error:
				if flags.verbose {
					panic(r)
				}

				log.ErrX(Exception, r.Error())
			case string:
				if flags.verbose {
					panic(r)
				}

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
