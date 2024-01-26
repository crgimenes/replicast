package luaengine

import (
	"sync"

	"replicast/config"

	lua "github.com/yuin/gopher-lua"
)

var (
	mx sync.Mutex
)

func ConfigDebug(L *lua.LState) int {
	mx.Lock()
	defer mx.Unlock()
	if L.GetTop() == 1 {
		config.CFG.Debug = L.ToBool(1)
	}
	L.Push(lua.LBool(config.CFG.Debug))
	return 1
}

func ConfigListen(L *lua.LState) int {
	mx.Lock()
	defer mx.Unlock()
	if L.GetTop() == 1 {
		config.CFG.Listen = L.ToString(1)
	}
	L.Push(lua.LString(config.CFG.Listen))
	return 1
}

func ConfigPath(L *lua.LState) int {
	mx.Lock()
	defer mx.Unlock()
	if L.GetTop() == 1 {
		config.CFG.Path = L.ToString(1)
	}
	L.Push(lua.LString(config.CFG.Path))
	return 1
}

func Startup(initLua string) error {
	L := lua.NewState()
	defer L.Close()
	L.SetGlobal("ConfigPath", L.NewFunction(ConfigPath))
	L.SetGlobal("ConfigDebug", L.NewFunction(ConfigDebug))
	L.SetGlobal("ConfigListen", L.NewFunction(ConfigListen))
	return L.DoFile(initLua)
}
