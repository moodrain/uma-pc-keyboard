package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type Point struct {
	X int
	Y int
}

var mouseX int
var mouseY int
var keyMap map[rune]Point
var ctrlMap map[uint16]bool

func main() {
	keyMap = make(map[rune]Point)
	ctrlMap = make(map[uint16]bool)
	stop := false
	EvChan := hook.Start()
	defer hook.End()
	fmt.Println("start listen, press ctrl + alt + p to stop/start")
	for ev := range EvChan {
		if ev.Kind == hook.MouseMove {
			mouseX = int(ev.X)
			mouseY = int(ev.Y)
		}
		code := ev.Keycode
		char := ev.Keychar
		if code != 0 {
			if ev.Kind == hook.KeyHold {
				ctrlMap[code] = true
			} else if ev.Kind == hook.KeyUp {
				ctrlMap[code] = false
			}
			if ctrlMap[25] && ctrlMap[29] && ctrlMap[56] {
				stop = !stop
				if stop {
					fmt.Println("stop")
				} else {
					fmt.Println("restart")
				}
			}
		} else {
			if ev.Kind == hook.KeyDown {
				if stop {
					continue
				}
				if _, ok := keyMap[char]; ! ok {
					fmt.Println("set new bind " + string(char))
					keyMap[char] = Point{
						X: mouseX,
						Y: mouseY,
					}
				} else {
					p := keyMap[char]
					robotgo.MoveMouse(p.X, p.Y)
					robotgo.MouseClick("left")
				}
			}
		}
	}
}