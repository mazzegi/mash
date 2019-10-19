package mash

import (
	"fmt"

	"github.com/fatih/color"
)

type Formatter interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

func DefaultFormatter() Formatter {
	return defaultFormatter{}
}

type defaultFormatter struct{}

func (f defaultFormatter) Infof(format string, args ...interface{}) {
	color.Cyan("INFO: %s", fmt.Sprintf(format, args...))
}

func (f defaultFormatter) Errorf(format string, args ...interface{}) {
	color.Red("ERROR: %s", fmt.Sprintf(format, args...))
}
