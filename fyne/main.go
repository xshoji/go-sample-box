package main

import (
//	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/app"
)

func main() {
	app := app.New()

	w := app.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Exit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}