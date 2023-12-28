package hook

import (
	"log/slog"
	"testing"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func TestDafaultHook(t *testing.T) {
	hook.Register(hook.MouseDown, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] {
			go func() {
				time.Sleep(100 * time.Millisecond)
				robotgo.KeyTap("c", "ctrl")
				queryText, _ := robotgo.ReadAll()
				slog.Info("queryText: ", queryText)
			}()
		}
		// HookChan <- struct{}{}
		// lastKeyPressTime = nil
	})

	hook.Register(hook.MouseHold, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] {
			lastKeyPressTime = time.Now()
		}
	})

	hook.Register(hook.MouseUp, []string{}, func(e hook.Event) {
		if e.Button == hook.MouseMap["center"] {
			robotgo.KeyTap("c", "ctrl")
			queryText, _ := robotgo.ReadAll()
			slog.Info("queryText: ", queryText)
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}
