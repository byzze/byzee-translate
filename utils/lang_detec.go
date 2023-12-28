package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"handy-translate/config"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

var langMap = map[string]string{
	"zh":  "zh_cn",
	"cht": "zh_tw",
	"en":  "en",
	"jp":  "ja",
	"kor": "ko",
	"fra": "fr",
	"spa": "es",
	"ru":  "ru",
	"de":  "de",
	"it":  "it",
	"tr":  "tr",
	"pt":  "pt_pt",
	"vie": "vi",
	"id":  "id",
	"th":  "th",
	"may": "ms",
	"ar":  "ar",
	"hi":  "hi",
	"nob": "nb_no",
	"nno": "nn_no",
	"per": "fa",
}

func baiduDetect(text string) (lang string) {
	lang = "en"

	apiURL := "https://fanyi.baidu.com/langdetect"
	data := url.Values{}
	data.Set("query", text)

	client := http.Client{}
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		slog.Error("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		slog.Error("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Error reading response body:", err)
		return
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return lang
	}

	if result["error"].(float64) == 0 {
		if lan, ok := result["lan"].(string); ok {
			if mappedLang, ok := langMap[lan]; ok {
				lang = mappedLang
				return
			}
		}
	}

	return
}

// LangDetect 检测查询文本语言类型
func LangDetect(text string) string {
	langDetectEngine := config.Data.TranslateWay

	switch langDetectEngine {
	default:
		return baiduDetect(text)
		// case "google":
		// 	return googleDetect(text), nil
		// case "local":
		// 	return localDetect(text), nil
		// case "tencent":
		// 	return tencentDetect(text), nil
		// case "niutrans":
		// 	return niutransDetect(text), nil
		// case "yandex":
		// 	return yandexDetect(text), nil
		// case "bing":
		// 	return bingDetect(text), nil
	}
}
