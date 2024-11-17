package command

import "github.com/go-vgo/robotgo"

func moveCursor(dx int, dy int) {
	// current cursor position
	x, y := robotgo.GetMousePos()
	robotgo.Move(x+dx, y+dy)
}
