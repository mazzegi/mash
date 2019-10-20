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
	formatter Formatter
	lstate    *lua.LState
}

func New(options ...Option) (*Mash, error) {
	m := &Mash{
		args:      []string{},
		formatter: DefaultFormatter(),
		lstate:    lua.NewState(),
	}
	for _, o := range options {
		err := o(m)
		if err != nil {
			return nil, err
		}
	}
	m.RegisterHandler("mash", NewCoreHandler(m.args, m.formatter))
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
	m.RegisterHandler("time", NewTimeHandler())
}

func (m *Mash) RunScript(script string) error {
	return m.lstate.DoString(script)
}

//CoreHandler is the handler, to handle basic mash commands
type CoreHandler struct {
	args      []string
	formatter Formatter
}

func NewCoreHandler(args []string, formatter Formatter) *CoreHandler {
	return &CoreHandler{
		args:      args,
		formatter: formatter,
	}
}

func (h *CoreHandler) Infof(format string, args ...interface{}) {
	h.formatter.Infof(format, args...)
}

func (h *CoreHandler) Errorf(format string, args ...interface{}) {
	h.formatter.Errorf(format, args...)
}

func (h *CoreHandler) Try(r Result, exitCode int) {
	if r.Ok() {
		return
	}
	h.formatter.Errorf("%s: %s", r.Context(), r.ErrorText())
	h.Exit(exitCode)
}

func (h *CoreHandler) Exit(code int) {
	os.Exit(code)
}

func (h *CoreHandler) NumArgs() int {
	return len(h.args)
}

func (h *CoreHandler) Arg(i int) string {
	if i < 0 || i >= len(h.args) {
		return ""
	}
	return h.args[i]
}

func (h *CoreHandler) Exec(prg string, args ...string) Result {
	cmd := exec.Command(prg, args...)
	bs, err := cmd.CombinedOutput()
	return NewResult(fmt.Sprintf("exec (%s)", prg), string(bs), err)
}
