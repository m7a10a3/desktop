// +build !linux ci

package main

import (
	"log"
	"runtime"

	"fyne.io/fyne"

	"github.com/fyne-io/desktop"
)

func setupDesktop(a fyne.App) desktop.Desktop {
	log.Println("Full desktop not possible on", runtime.GOOS)
	return desktop.NewEmbeddedDesktop(a)
}
