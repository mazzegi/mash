package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/mazzegi/mash"
	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) < 2 {
		errorf("to less arguments for mash. Usage: 'mash script.lua'")
		os.Exit(1)
	}
	script, err := prepareScript(os.Args[1])
	if err != nil {
		errorf("prepare script: %v", err)
		os.Exit(2)
	}
	m, err := mash.New(script, mash.WithArgs(os.Args[2:]...))
	if err != nil {
		errorf("new mash: %v", err)
		os.Exit(3)
	}
	if err := m.Run(); err != nil {
		errorf("run-mash: %v", err)
		os.Exit(4)
	}
}

func prepareScript(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", errors.Wrapf(err, "open-file (%s)", file)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "#!") {
			// removecomment out shebang line
			line = "-- " + line
		}
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return "", errors.Errorf("error scanning script-file : %v", scanner.Err())
	}

	return strings.Join(lines, "\n"), nil
}

func errorf(format string, args ...interface{}) {
	mash.DefaultFormatter().Errorf(format, args...)
}
