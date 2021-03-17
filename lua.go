package main

import (
	"fmt"

	"github.com/yuin/gopher-lua"
)

var l *lua.LState
var lcart *lua.LState

func InitLua() {
	l = lua.NewState()
	lcart = lua.NewState()

	l.OpenLibs()
	lcart.OpenLibs()

	l.SetGlobal("dofile", l.NewFunction(m9dofile))
	lcart.SetGlobal("dofile", l.NewFunction(m9empty))

	SetGlobalF("debugprint", l.NewFunction(m9debugprint))
	SetGlobalF("print", l.NewFunction(m9print))

	l.DoString("dofile('hello.lua')")
}

func SetGlobalF(name string, fun lua.LValue) {
	l.SetGlobal(name, fun)
	lcart.SetGlobal(name, fun)
}

func m9dofile(l *lua.LState) int {
	cart := l.ToString(1)
	err := lcart.DoFile(cart)
	if err != nil {
		panic(err)
	}

	return 1
}

func m9empty(l *lua.LState) int {
	return 1
}

func m9debugprint(l *lua.LState) int {
	text := l.ToString(1)
	fmt.Println(text)

	return 1
}

func m9print(L *lua.LState) int {
	text := L.ToString(1)
	x := L.ToInt(2)
	y := L.ToInt(3)
	textcolor := L.ToInt(4)

	c := palette[textcolor]
	xx := x
	sx := x
	yy := y

	for _, ch := range text {
		char := font[string(ch)]
		for i := 0; i < char.Height; i++ {
		 	bin := char.Data[i]

			for _, pix := range bin {
				if string(pix) == "1" {
					setpixel(
						xx + sx,
						char.Y + yy,
						c)
				}
			 	xx += 1
			}
			yy += 1
			xx = x
		}
		sx += char.Width
		yy = y
	}

	return 1
}
