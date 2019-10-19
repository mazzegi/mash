package mash

import (
	"os"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type Option func(m *Mash) error

func WithArgs(args ...string) Option {
	return func(m *Mash) error {
		m.args = args
		return nil
	}
}

func WithFormatter(f Formatter) Option {
	return func(m *Mash) error {
		m.formatter = f
		return nil
	}
}

type Mash struct {
	args      []string
	script    string
	formatter Formatter
	lstate    *lua.LState
}

func New(script string, options ...Option) (*Mash, error) {
	m := &Mash{
		args:      []string{},
		script:    script,
		formatter: DefaultFormatter(),
		lstate:    lua.NewState(),
	}
	for _, o := range options {
		err := o(m)
		if err != nil {
			return nil, err
		}
	}
	m.lstate.SetGlobal("mash", luar.New(m.lstate, m))
	m.lstate.SetGlobal("print", m.lstate.NewFunction(m.formatter.luaPrint))
	m.lstate.SetGlobal("print_error", m.lstate.NewFunction(m.formatter.luaPrintError))
	return m, nil
}

func (m *Mash) Run() error {
	return m.lstate.DoString(m.script)
}

func (m *Mash) Try(r Result, exitCode int) {
	if !r.IsError() {
		return
	}
	m.formatter.Errorf("%s: %s", r.Context(), r.ErrorText())
	m.Exit(exitCode)
}

func (m *Mash) Exit(code int) {
	os.Exit(code)
}

func (m *Mash) NumArgs() int {
	return len(m.args)
}

func (m *Mash) Arg(i int) string {
	if i < 0 || i >= len(m.args) {
		return ""
	}
	return m.args[i]
}
