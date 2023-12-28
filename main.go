package main

import (
	"embed"
	_ "embed"
	"fmt"
	"handy-translate/config"
	"handy-translate/hook"
	"handy-translate/screenshot"
	"handy-translate/toolbar"
	"handy-translate/translate"
	"handy-translate/translate/baidu"
	"handy-translate/translate/youdao"
	"log"
	"log/slog"
	"reflect"
	"regexp"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed frontend/dist
var assets embed.FS

//go:embed frontend/public/appicon.png
var iconlogo []byte

var app *application.App

var appInfo = &App{}

var fromLang, toLang = "auto", "zh"

var projectName = "handy-translate"

func main() {
	app = application.New(application.Options{
		Name: projectName,
		Bind: []any{
			&App{},
		},
		Icon: iconlogo,
		Assets: application.AssetOptions{
			FS: assets,
		},
	})

	toolbar.NewWindow(app)

	translate.NewWindow(app)

	screenshot.NewWindow(app)

	app.Events.On("translateLang", func(event *application.WailsEvent) {
		app.Logger.Info("translateType", slog.Any("event", event))

		valueType := reflect.TypeOf(event.Data)
		fmt.Println("Type:", valueType)

		if optionalData, ok := event.Data.([]interface{}); ok {
			fromLang = fmt.Sprintf("%v", optionalData[0])
			toLang = fmt.Sprintf("%v", optionalData[1])
			app.Logger.Info("translateLang",
				slog.String("fromLang", fromLang),
				slog.String("toLang", toLang))
		}
	})

	// 系统托盘
	systemTray := app.NewSystemTray()
	myMenu := app.NewMenu()

	myMenu.Add("翻译").OnClick(func(ctx *application.Context) {
		translate.Window.Center()
		translate.Window.Show()
	})

	myMenu.Add("截图").OnClick(func(ctx *application.Context) {
		screenshot.ScreenshotFullScreen()
		base64Image := screenshot.ScreenshotFullScreen()
		app.Events.Emit(&application.WailsEvent{Name: "screenshotBase64", Data: base64Image})
	})

	myMenu.Add("退出").OnClick(func(ctx *application.Context) {
		app.Quit()
	})

	systemTray.SetMenu(myMenu)
	systemTray.SetIcon(iconlogo)

	systemTray.OnClick(func() {
		toolbar.Window.Show()
	})

	// 初始化文件和鼠标事件
	config.Init(projectName)
	go processHook()

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func sendDataToJS(query, result, explains string) {
	sendQueryText(query)
	sendResult(result, explains)
}

func sendQueryText(queryText string) {
	app.Events.Emit(&application.WailsEvent{Name: "query", Data: queryText})
}

func sendResult(result, explains string) {
	app.Events.Emit(&application.WailsEvent{Name: "result", Data: result})
	app.Events.Emit(&application.WailsEvent{Name: "explains", Data: explains})
}

func isChinese(text string) bool {
	// 使用正则表达式匹配中文字符
	re := regexp.MustCompile("[\u4e00-\u9fa5]")
	return re.MatchString(text)
}

// 监听处理鼠标事件
func processHook() {
	go hook.DafaultHook(app) // 使用robotgo处理

	for {
		select {
		case <-hook.HookChan:
			queryText, _ := robotgo.ReadAll()

			// toolBar 中英互译
			if isChinese(queryText) {
				toLang = "en"
			}

			if !isChinese(queryText) {
				switch config.Data.TranslateWay {
				case youdao.Way:
					toLang = "zh-CHS"
				case baidu.Way:
					toLang = "zh"
				}
			}

			app.Logger.Info("GetQueryText",
				slog.String("queryText", queryText),
				slog.String("fromLang", fromLang),
				slog.String("toLang", toLang))

			if queryText != translate.GetQueryText() && queryText != "" {
				translate.SetQueryText(queryText)
				translateRes := processTranslate(queryText)
				// 发送结果至前端
				if len(translateRes) == 0 {
					translateRes = queryText
				}
				sendDataToJS(queryText, translateRes, "")
				continue
			}

			processToolbarShow()
		}
	}
}

// 翻译处理
func processTranslate(queryText string) string {
	translateWay := translate.GetTransalteWay(config.Data.TranslateWay)

	result, err := translateWay.PostQuery(queryText, fromLang, toLang)
	if err != nil {
		slog.Error("PostQuery", err)
	}

	app.Logger.Info("Transalte",
		slog.Any("fromLang", fromLang),
		slog.Any("toLang", toLang),
		slog.Any("result", result),
		slog.Any("translateWay", translateWay.GetName()))

	translateRes := strings.Join(result, "\n")

	return translateRes
}
