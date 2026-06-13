package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tab Demo")

	tabs := container.NewAppTabs(
		container.NewTabItem("节点", widget.NewLabel("节点页面")),
		container.NewTabItem("规则", widget.NewLabel("规则页面")),
		container.NewTabItem("设置", widget.NewLabel("设置页面")),
	)

	w.SetContent(tabs)
	w.Resize(fyne.NewSize(400, 300))

	w.ShowAndRun()
}
