package main

import (
	"fmt"

	"github.com/yuin/gopher-lua"
)

func m9setfunc(L *lua.LState) int {
	funcname := L.ToString(1)
	function := L.ToFunction(2)
	ring := L.ToInt(3)

	switch ring {
	case 0:
		SetGlobalF(funcname, function)
	}

	return 0
}

func m9dofile(l *lua.LState) int {
	cart := l.ToString(1)
	started = true
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

func m9plot(L *lua.LState) int {
	x := L.ToInt(1)
	y := L.ToInt(2)
	color := L.ToInt(3)

	c := palette[color]
	setpixel(x, y, c)

	return 0
}

func m9button(L *lua.LState) int {
	key := L.ToInt(1)

	btn := buttons[key]
	if btn == true {
		lcart.Push(lua.LBool(true))
		buttons[key] = false
	} else {
		lcart.Push(lua.LBool(false))
	}

	return 1
}
