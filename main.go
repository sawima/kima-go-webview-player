package main

// include gtk header, to access the GTK related C-Types

/*
#cgo linux openbsd freebsd pkg-config: gtk+-3.0 webkit2gtk-4.0
#include <gtk/gtk.h>
*/
import "C"

import (
	"os"
	"os/signal"
	"time"

	"kima-go-webview-player/socketclient"

	"github.com/webview/webview"
)

var websites []string = []string{
	"https//kimacloud.online",
	"https//kimacloud.com.cn",
	"https://baidu.com",
	"https://taobao.com",
	"https://qq.com",
	"https://jd.com",
	"https://bi.aliyun.com/token3rd/screen/view/pc.htm?validityTime=120&bizdate=2023-07-11&pageId=0826bc66-066f-4805-a1b8-7e3cd230f275&accessId=INNER_60ee294786b8409aae1ed598e7e693bc&timestamp=1689061811298&signature=hImMl86Nm6NhsMQvcnpk6AgM6xftUJH4fXQQfIPhaBM%3D",
	"https://bi.aliyun.com/token3rd/screen/view/pc.htm?validityTime=120&bizdate=2023-07-11&pageId=6a80a6f0-268b-41af-8587-968794b046f9&accessId=INNER_60ee294786b8409aae1ed598e7e693bc&timestamp=1689061811298&signature=9uadF8AP7cT3cC2Vccdnwq%2BLIWzLgPrYUpKCxxp6bbM%3D",
	"https://bi.aliyun.com/token3rd/screen/view/pc.htm?validityTime=120&bizdate=2023-07-11&pageId=9090b08e-b9ce-4f1c-a854-e2bef8722522&accessId=INNER_60ee294786b8409aae1ed598e7e693bc&timestamp=1689061811298&signature=jiCVJsF0ma029HQUtpDxxvJ73cy7JFgO3Q1170ImG3A%3D",
}

const basicSite = "http://localhost:8081"

var navUrl = basicSite
var w = webview.New(true)

func main() {
	window := w.Window()

	defer w.Destroy()

	C.gtk_window_fullscreen((*C.GtkWindow)(window))

	w.SetTitle("Vending Machine")
	w.SetSize(480, 800, webview.HintNone)
	w.Navigate(navUrl)
	intervalFunc()
	w.Run()
}

func intervalFunc() {
	ticker := time.NewTicker(1 * time.Second)
	newUrl := ""
	reload := false
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		for {
			select {
			case <-ticker.C:
				newUrl, reload = socketclient.ReadSocket("fetch url")
				if reload {
					w.Navigate(basicSite)
				} else {
					if newUrl != navUrl {
						print("new url")
						navUrl = newUrl
						w.Navigate(newUrl)
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
