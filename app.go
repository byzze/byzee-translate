package main

import (
	"context"
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/translate"
	"handy-translate/translate/youdao"
	"handy-translate/utils"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

func (a *App) MyFetch(URL string, content map[string]interface{}) interface{} {
	return utils.MyFetch(URL, content)
}

func (a *App) sendQueryText(queryText string) {
	runtime.EventsEmit(a.ctx, "query", queryText)
}

func (a *App) sendResult(result, explian string) {
	runtime.EventsEmit(a.ctx, "result", result)
	runtime.EventsEmit(a.ctx, "explian", explian)
}

func (a *App) SendDataToJS(query, result, explian string) {
	logrus.WithFields(logrus.Fields{
		"query":   query,
		"result":  result,
		"explian": explian,
	}).Info("SendDataToJS", query, result, explian)

	a.sendQueryText(query)
	a.sendResult(result, explian)

}

// test data
func (a *App) onDomReady(ctx context.Context) {
	a.sendQueryText("启动成功")
	// system tray 系统托盘
	onReady := func() {
		systray.SetIcon(appicon)
		systray.SetTitle(config.Data.Appname)
		systray.SetTooltip(config.Data.Appname + "便捷翻译工具")
		mShow := systray.AddMenuItem("显示", "显示翻译工具")
		mQuitOrig := systray.AddMenuItem("退出", "退出翻译工具")
		// Sets the icon of a menu item. Only available on Mac and Windows.
		mShow.SetIcon(appicon)
		for {
			select {
			case <-mShow.ClickedCh:
				a.Show()
			case <-mQuitOrig.ClickedCh:
				a.Quit()
			}
		}
	}
	systray.Run(onReady, func() { logrus.Info("app quit") })
}

var fromLang, toLang = "auto", "zh"

func eventFunc(ctc context.Context) {
	runtime.EventsOn(ctc, "translateType", func(optionalData ...interface{}) {
		logrus.WithField("optionalData", optionalData).Info("translateType")
		if len(optionalData) >= 2 {
			fromLang = fmt.Sprintf("%v", optionalData[0])
			toLang = fmt.Sprintf("%v", optionalData[1])
		}
	})
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	config.Init(ctx)

	go hook.DafaultHook()
	go hook.WindowsHook()

	eventFunc(ctx)
	// scList, _ := runtime.ScreenGetAll(ctx)

	// var screenX, screenY int
	// for _, v := range scList {
	// 	if v.IsCurrent {
	// 		screenX = v.Width
	// 		screenY = v.Height
	// 	}
	// }

	runtime.WindowCenter(ctx)
	go func() {
		for {
			select {
			case <-hook.HookChan:
				logrus.Info("HookChan Process")
				// windowX, windowY := runtime.WindowGetSize(ctx)
				// x, y := robotgo.GetMousePos()
				// x, y = x+10, y-10
				// i := 0
				// bounds := screenshot.GetDisplayBounds(i)

				// img, err := screenshot.CaptureRect(bounds)
				// if err != nil {
				// 	panic(err)
				// }

				// tn := time.Now().UnixNano()
				// frontendName := fmt.Sprintf("screenshot/screenshot-%d.png", tn)
				// fileName := fmt.Sprintf("./frontend/%s", frontendName)
				// if _, err := os.Stat(fileName); err == nil {
				// 	// 文件存在，删除它
				// 	err := os.Remove(fileName)
				// 	if err != nil {
				// 		// 处理删除文件时的错误
				// 		logrus.WithError(err).Error("os.Remove")
				// 	}
				// 	println("文件已删除")
				// }

				// file, _ := os.Create(fileName)
				// defer file.Close()
				// png.Encode(file, img)

				// runtime.EventsEmit(a.ctx, "screenshot", frontendName)
				runtime.WindowShow(ctx)

				// queryText, _ := runtime.ClipboardGetText(a.ctx)

				// a.sendQueryText(queryText)

				// if queryText != hook.GetQueryText() {
				// 	fmt.Println("GetQueryText================", fromLang, toLang)
				// 	a.Transalte(queryText, fromLang, toLang)
				// }
				// TODO 弹出窗口根据鼠标位置变动
				// fmt.Println("or:", x, y, screenX, screenY, windowX, windowY)
				// if y+windowY+20 >= screenY {
				// 	y = screenY - windowY - 20
				// }

				// if x+windowX >= screenX {
				// 	x = screenX - windowX
				// }
				// fmt.Println("new:", x, y, screenX, screenY, windowX, windowY)
				// runtime.WindowSetPosition(ctx, x, y)
			}
		}
	}()

}

// Greet returns a greeting for the given name
func (a *App) GetKeyBoard() []string {
	if len(config.Data.Keyboard) == 0 {
		config.Data.Keyboard = make([]string, 3)
	}
	return config.Data.Keyboard
}

func (a *App) SetKeyBoard(ctrl, shift, key string) {
	config.Data.Keyboard = []string{ctrl, shift, key}
	logrus.Info(config.Data.Keyboard)
	config.Save()
	go hook.Hook()
}

func (a *App) GetTransalteMap() string {
	var translateList = config.Data.Translate
	bTranslate, err := json.Marshal(translateList)
	if err != nil {
		logrus.WithError(err).Error("Marshal")
	}
	return string(bTranslate)
}

func (a *App) SetTransalteWay(translateWay string) {
	fmt.Println(translateWay)
	config.Data.TranslateWay = translateWay
	hook.SetQueryText("")
	config.Save()
	logrus.WithField("config.Data.Translate", config.Data.Translate).Info("SetTransalteList")
}

func (a *App) GetTransalteWay() string {
	return config.Data.TranslateWay
}

func (a *App) Transalte(queryText, fromLang, toLang string) {
	hook.SetQueryText(queryText)
	// 加载动画loading
	runtime.EventsEmit(a.ctx, "loading", "true")
	time.Sleep(time.Second * 3)
	defer runtime.EventsEmit(a.ctx, "loading", "false")

	transalteWay := translate.GetTransalteWay(config.Data.TranslateWay)

	logrus.WithFields(logrus.Fields{
		"queryText":    queryText,
		"transalteWay": transalteWay.GetName(),
		"fromLang":     fromLang,
		"toLang":       toLang,
	}).Info("Transalte")

	curName := transalteWay.GetName()
	// 使用 strings.Replace 替换 \r 和 \n 为空格
	queryTextTmp := strings.ReplaceAll(queryText, "\r", "")
	queryTextTmp = strings.ReplaceAll(queryTextTmp, "\n", "")

	result, err := transalteWay.PostQuery(queryTextTmp, fromLang, toLang)
	if err != nil {
		logrus.WithError(err).Error("PostQuery")
	}

	logrus.WithFields(logrus.Fields{
		"result": result,
	}).Info("Transalte")

	if len(result) >= 2 && curName == youdao.Way {
		a.SendDataToJS(queryText, result[0], result[1])
	}

	transalteRes := strings.Join(result, ",")
	a.SendDataToJS(queryText, transalteRes, "")

}

func (a *App) Quit() {
	runtime.Quit(a.ctx)
	systray.Quit()
}

func (a *App) Show() {
	runtime.WindowShow(a.ctx)
}
