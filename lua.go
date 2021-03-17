package main

import (
	"fmt"

	"github.com/yuin/gopher-lua"
)

var ring int
var l *lua.LState

func InitLua() {
	ring = 0
	l = lua.NewState()

	l.OpenLibs()

	l.SetGlobal("dofile", l.NewFunction(m9dofile))
	l.SetGlobal("debugprint", l.NewFunction(m9debugprint))

	l.DoString("dofile('hello.lua')")
}

func m9dofile(l *lua.LState) int {
	ring += 1
	if ring == 0 {

		cart := l.ToString(1)
		err := l.DoFile(cart)
		if err != nil {
			panic(err)
		}
	}

	return 1
}

func m9debugprint(l *lua.LState) int {
	text := l.ToString(1)
	fmt.Println(text)

	return 1
}
