package main

import (
	"fmt"
	"io/ioutil"

	"github.com/riftbit/go-systray"
	"github.com/webview/webview"
)

func Show() {
	go func() {
		w := webview.New(false)
		defer w.Destroy()
		w.SetTitle("data center")
		w.SetSize(1062, 960, webview.HintNone)
		//w.SetHtml("Thanks for using webview!")
		w.Navigate("https://www.baidu.com")
		w.Run()
	}()
}
func onReady() {
	ico, err := ioutil.ReadFile("favicon.ico")
	if err == nil {
		systray.SetIcon(ico)
	}
	submenu := systray.AddMenuItem("显示", "显示主页", 0)

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出应用", 0)

	go func() {
		for {
			select {
			case <-submenu.OnClickCh():
				Show()
			case <-mQuit.OnClickCh():
				systray.Quit()
				return
			}
		}
	}()
}

func onExit() {
	// Cleaning stuff here.
}
func main() {
	go Show()
	systray.Run(onReady, onExit)
}
