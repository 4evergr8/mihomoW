package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("demo")

	w.SetContent(widget.NewButton("点击", func() {
		w.SetTitle("clicked")
	}))

	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
