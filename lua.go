package main

import (
	"github.com/yuin/gopher-lua"
)

var l *lua.LState
var lcart *lua.LState
var started = false

func InitLua() {
	CatchControls()

	l = lua.NewState()
	lcart = lua.NewState()

	l.OpenLibs()
	lcart.OpenLibs()

	l.SetGlobal("dofile", l.NewFunction(m9dofile))
	l.SetGlobal("setfunc", l.NewFunction(m9setfunc))
	lcart.SetGlobal("dofile", l.NewFunction(m9empty))

	SetGlobalF("debugprint", l.NewFunction(m9debugprint))
	SetGlobalF("print", l.NewFunction(m9print))
	SetGlobalF("plot", l.NewFunction(m9plot))
	SetGlobalF("button", l.NewFunction(m9button))

	if err := l.DoFile("mocha.lua"); err != nil {
		panic(err)
	}
}

func SetGlobalF(name string, fun lua.LValue) {
	l.SetGlobal(name, fun)
	lcart.SetGlobal(name, fun)
}

