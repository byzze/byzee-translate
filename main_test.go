package main

import (
	_ "embed"
	"fmt"
	"handy-translate/os_api/windows"
	"log/slog"
	"syscall"
	"testing"

	"github.com/go-vgo/robotgo"
	"github.com/lxn/win"
	hook "github.com/robotn/gohook"
)

// GetCursorPos 获取鼠标位置 github.com/lxn/win
func GetCursorPos() *win.POINT {
	lpPoint := &win.POINT{}
	win.GetCursorPos(lpPoint)
	return lpPoint

}

func TestMouseClickPos(t *testing.T) {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] {
			x, y := robotgo.Location()
			fmt.Printf("Location:[x:%d,y:%d]\n", x, y)
			pos := GetCursorPos()
			fmt.Printf("GetCursorPos[x:%d,y:%d]\n", pos.X, pos.Y)
		}
	})
	s := hook.Start()
	<-hook.Process(s)
}

func TestWindowshow(t *testing.T) {
	windowName := "ToolBar"
	w := windows.FindWindow(windowName)
	w.ShowForWindows()

	lpWindowName, err := syscall.UTF16PtrFromString(windowName)
	if err != nil {
		slog.Error("UTF16PtrFromString", err)
	}

	// find window
	hwnd := win.FindWindow(nil, lpWindowName)
	if hwnd == 0 {
		slog.Error("FindWindow Failed")
	}
}
