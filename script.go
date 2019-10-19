package mash

import (
	"bufio"
	"os"
	"strings"

	"github.com/pkg/errors"
)

func LoadScript(file string) (string, error) {
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
			//remove comment out shebang line
			line = "-- " + line
		}
		lines = append(lines, line)
	}
	if scanner.Err() != nil {
		return "", errors.Errorf("error scanning script-file : %v", scanner.Err())
	}

	return strings.Join(lines, "\n"), nil
}
