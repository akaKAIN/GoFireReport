package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
)

func SignalCloseWindow(win *gtk.Window) error {
	// Инициализация сигнала закрытия окна
	_, err := win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	if err != nil {
		return fmt.Errorf("Инициализация сигнала закрытия окна: %s", err)
	}
	return nil
}

