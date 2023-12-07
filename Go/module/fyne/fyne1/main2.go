package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

func main() {
	a := app.New()
	window := a.NewWindow("clock")
	clock := widget.NewLabel("")
	updateTime(clock)
	window.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	window.ShowAndRun()
}

func updateTime(clock *widget.Label) {
	clock.SetText(time.Now().Format(time.DateTime))
}
