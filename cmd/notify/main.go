package main

import (
	gosxnotifier "github.com/CuiYao631/gosx-notifier"
	"log"
)

func main() {
	note := gosxnotifier.NewNotification("Check your Apple Stock!")

	//(可选) 设置标题
	note.Title = "title"

	//(可选) 设置字幕
	note.Subtitle = "message"

	//(可选) 从预定义的集合中设置声音.
	note.Sound = gosxnotifier.Default

	//(可选) 设置一个组，确保只显示一个通知，替换同一组id的先前通知.
	note.Group = "com.wails.DevMate"

	//(可选) 设置发件人 (通知现在将使用Safari图标)
	note.Sender = "com.wails.DevMate"

	//(可选) 指定要在单击通知时打开的url或bundleid。
	note.Link = "http://finance.yahoo.com/q?s=AAPL" //or BundleID like: com.apple.Terminal

	//(可选) 应用程序图标 (仅限10.9)
	//note.AppIcon = "build/windows/icon.ico"

	//(可选) 内容图像 (仅限10.9)
	note.ContentImage = "build/appicon.png"

	//然后，推送通知
	err := note.Push()

	//检查错误
	if err != nil {
		log.Println("Uh oh!")
	}
}
