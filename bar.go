package desktop

import (
	"os/exec"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func clockTick(clock, date *widget.Label) {
	tick := time.NewTicker(time.Second)
	go func() {
		for {
			<-tick.C
			clock.SetText(formattedTime())
			date.SetText(formattedDate())
		}
	}()
}

func formattedTime() string {
	return time.Now().Format("15:04:05")
}

func formattedDate() string {
	return time.Now().Format("2 Jan")
}

func createClock() *widget.Box {
	var style fyne.TextStyle
	style.Monospace = true

	clock := &widget.Label{
		Text:      formattedTime(),
		Alignment: fyne.TextAlignCenter,
		TextStyle: style,
	}
	date := &widget.Label{
		Text:      formattedDate(),
		Alignment: fyne.TextAlignCenter,
		TextStyle: style,
	}

	go clockTick(clock, date)

	return widget.NewHBox(clock, date)
}

func (l *deskLayout) newBar() fyne.CanvasObject {
	quit := widget.NewButtonWithIcon("", theme.CancelIcon(), func() {
		l.Root().Close()
	})
	clock := createClock()
	buttons := fyne.NewContainerWithLayout(layout.NewGridLayout(5),
		widget.NewButton("Browser", func() {
			exec.Command("chromium").Start()
		}),
		widget.NewButton("Terminal", func() {
			exec.Command("xterm").Start()
		}),
	)

	content := fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, quit, clock),
		quit,
		clock,
		buttons,
	)
	return fyne.NewContainerWithLayout(layout.NewMaxLayout(),
		canvas.NewRectangle(theme.BackgroundColor()),
		content,
	)
}
