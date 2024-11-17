package command

import "github.com/go-vgo/robotgo"

func handleMouseClick(button string) {
	// current cursor position
	robotgo.Click(button)
}
