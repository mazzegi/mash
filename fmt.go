package mash

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	lua "github.com/yuin/gopher-lua"
)

type Formatter interface {
	luaPrint(l *lua.LState) int
	luaPrintError(l *lua.LState) int
	LuaInfof(format string, args ...interface{})
	LuaErrorf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

func DefaultFormatter() Formatter {
	return defaultFormatter{}
}

type defaultFormatter struct{}

func (f defaultFormatter) LuaInfof(format string, args ...interface{}) {
	color.Cyan("LUA-INFO: %s", fmt.Sprintf(format, args...))
}

func (f defaultFormatter) LuaErrorf(format string, args ...interface{}) {
	color.Red("LUA-ERROR: %s", fmt.Sprintf(format, args...))
}

func (f defaultFormatter) Infof(format string, args ...interface{}) {
	color.Cyan("INFO: %s", fmt.Sprintf(format, args...))
}

func (f defaultFormatter) Errorf(format string, args ...interface{}) {
	color.Red("ERROR: %s", fmt.Sprintf(format, args...))
}

func luaCollectStringArgs(l *lua.LState) []string {
	num := l.GetTop()
	sl := make([]string, num)
	for i := 0; i < num; i++ {
		lv := l.Get(i + 1)
		sl[i] = lv.String()
	}
	return sl
}

func (f defaultFormatter) luaPrint(l *lua.LState) int {
	f.LuaInfof(strings.Join(luaCollectStringArgs(l), " "))
	return 0
}

func (f defaultFormatter) luaPrintError(l *lua.LState) int {
	f.LuaErrorf(strings.Join(luaCollectStringArgs(l), " "))
	return 0
}
