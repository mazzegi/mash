package mash

import lua "github.com/yuin/gopher-lua"

type LuaFunc func(args ...string) Result

type LuaFuncMap map[string]LuaFunc

func LuaStringArgs(l *lua.LState) []string {
	num := l.GetTop()
	sl := make([]string, num)
	for i := 0; i < num; i++ {
		lv := l.Get(i + 1)
		sl[i] = lv.String()
	}
	return sl
}

func LuaLGFunc(f LuaFunc) lua.LGFunction {
	return func(l *lua.LState) int {
		res := f(LuaStringArgs(l)...)
		l.Push(lua.LString(res.Value()))
		return 1
	}
}
