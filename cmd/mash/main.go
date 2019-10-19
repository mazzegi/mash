package main

import (
	"os"
	"strings"

	"github.com/mazzegi/mash"
)

func main() {
	if len(os.Args) < 2 {
		errorf("to less arguments for mash. Usage: 'mash script.lua'")
		os.Exit(1)
	}
	script, err := mash.LoadScript(os.Args[1])
	if err != nil {
		errorf("prepare script: %v", err)
		os.Exit(2)
	}
	m, err := mash.New(script, mash.WithArgs(os.Args[2:]...))
	if err != nil {
		errorf("new mash: %v", err)
		os.Exit(3)
	}

	//register all default handler
	m.RegisterDefaultHandler()

	//overwrite lua:print
	m.RegisterLuaFunc("print", func(args ...string) mash.Result {
		return infof(strings.Join(args, " "))
	})
	//register print_error
	m.RegisterLuaFunc("print_error", func(args ...string) mash.Result {
		return errorf(strings.Join(args, " "))
	})

	if err := m.Run(); err != nil {
		errorf("run-mash: %v", err)
		os.Exit(4)
	}
}

func infof(format string, args ...interface{}) mash.Result {
	mash.DefaultFormatter().Infof(format, args...)
	return mash.NoResult
}

func errorf(format string, args ...interface{}) mash.Result {
	mash.DefaultFormatter().Errorf(format, args...)
	return mash.NoResult
}
