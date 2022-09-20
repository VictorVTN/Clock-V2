package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	for true {
		timeNow := time.Now().Format(time.ANSIC)
		clock.SetText(timeNow)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Relogio V2")

	clock := widget.NewLabel("")
	clock.Alignment = fyne.TextAlignCenter
	w.SetContent(container.NewVBox(
		clock,
		widget.NewButton("QUERO SABER O HORARIO", func() {
			go updateTime(clock)
		}),
	))

	w.ShowAndRun()
}
