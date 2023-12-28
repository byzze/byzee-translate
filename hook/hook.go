package hook

import (
	"handy-translate/config"
	"handy-translate/os_api/windows"
	"handy-translate/screenshot"
	"log/slog"
	"runtime"

	"sync"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// HookChan channle
var HookChan = make(chan struct{}, 1)

var defaulthook = func(e hook.Event) {
	if e.Button == hook.MouseMap["center"] {
		slog.Info("mouse center down", slog.Any("evnet", e))
		robotgo.KeyTap("c", "ctrl")
		HookChan <- struct{}{}
	}
}

var keyboardhook = func(e hook.Event) {
	if pressLock.TryLock() {
		slog.Info("keyboardhook", slog.Any("event", e))

		robotgo.KeyTap("c", "ctrl")
		time.Sleep(time.Millisecond * 100)
		HookChan <- struct{}{}
		pressLock.Unlock()
	}
}

var lastKeyPressTime time.Time

var lastMouseTime time.Time

var pressCount int

// DafaultHook register hook event
func DafaultHook(app *application.App) {
	if runtime.GOOS == "windows" {
		go windows.WindowsHook(HookChan) // 完善，robotgo处理的不完美, 使用windows 原生api
	} else {
		// default mid mouse
		hook.Register(hook.MouseDown, []string{}, defaulthook)
	}

	hook.Register(hook.KeyDown, []string{"c", "ctrl"}, func(e hook.Event) {
		pressLock.Lock()
		defer pressLock.Unlock()
		if pressCount == 0 {
			lastKeyPressTime = time.Now()
			pressCount++
		} else {
			elapsed := time.Since(lastKeyPressTime)
			// Check if the time elapsed is greater than 500 milliseconds
			if elapsed.Milliseconds() < 300 {
				slog.Info("ctrl+c+c", slog.Any("event", e))
				HookChan <- struct{}{}
				// Reset pressCount after successful trigger
			}
			pressCount = 0
		}
	})

	screenshotKey := config.Data.Keyboards["screenshot"]
	hook.Register(hook.KeyDown, screenshotKey, func(e hook.Event) {
		base64Image := screenshot.ScreenshotFullScreen()
		app.Events.Emit(&application.WailsEvent{Name: "screenshotBase64", Data: base64Image})

	})

	s := hook.Start()
	<-hook.Process(s)
}

var pressLock sync.RWMutex

// ToolBarHook register hook event 用于配置快捷键 TODO
func ToolBarHook() {
	slog.Info("--- Please wait hook starting ---")
	hook.End()
	if len(config.Data.Keyboards) == 0 || config.Data.Keyboards["center"][0] == "center" {
		hook.Register(hook.MouseDown, []string{}, defaulthook)
	} else {
		hook.Register(hook.KeyDown, config.Data.Keyboards["toolBar"], keyboardhook)
	}

	s := hook.Start()
	<-hook.Process(s)
}
