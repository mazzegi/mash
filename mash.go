package mash

import (
	"fmt"
	"os"
	"os/exec"

	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type Handler interface{}

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
	//register this as lua-global
	m.lstate.SetGlobal("mash", luar.New(m.lstate, m))
	return m, nil
}

func (m *Mash) RegisterLuaFunc(name string, fn LuaFunc) {
	m.lstate.SetGlobal(name, m.lstate.NewFunction(LuaLGFunc(fn)))
}

func (m *Mash) RegisterLuaFuncMap(fm LuaFuncMap) {
	for name, fn := range fm {
		m.RegisterLuaFunc(name, fn)
	}
}

func (m *Mash) RegisterHandler(name string, h Handler) {
	m.lstate.SetGlobal(name, luar.New(m.lstate, h))
}

func (m *Mash) RegisterDefaultHandler() {
	m.RegisterHandler("fs", NewFSHandler())
}

func (m *Mash) Infof(format string, args ...interface{}) {
	m.formatter.Infof(format, args...)
}

func (m *Mash) Errorf(format string, args ...interface{}) {
	m.formatter.Errorf(format, args...)
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

func (m *Mash) Exec(prg string, args ...string) Result {
	cmd := exec.Command(prg, args...)
	bs, err := cmd.CombinedOutput()
	return NewResult(fmt.Sprintf("exec (%s)", prg), string(bs), err)
}
