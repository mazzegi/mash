package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mazzegi/mash"
)

func main() {
	var interactive = false
	for _, arg := range os.Args {
		if arg == "-i" {
			interactive = true
			break
		}
	}

	m, err := mash.New(mash.WithArgs(os.Args...))
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

	if !interactive {
		if len(os.Args) < 2 {
			errorf("missing script argument for non-interactive (batch) mode")
			os.Exit(1)
		}
		script, err := mash.LoadScript(os.Args[1])
		if err != nil {
			errorf("prepare script: %v", err)
			os.Exit(2)
		}
		if err := m.RunScript(script); err != nil {
			errorf("run-mash: %v", err)
			os.Exit(4)
		}
	} else {
		prompt := func() { fmt.Printf("> ") }
		scanner := bufio.NewScanner(os.Stdin)
		prompt()
		for scanner.Scan() {
			s := scanner.Text()
			if err := m.RunScript(s); err != nil {
				errorf("run-mash: %v", err)
			}
			prompt()
		}
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
