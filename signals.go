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
		return fmt.Errorf("Инициализация сигнала закрытия окна: %s\n", err)
	}
	return nil
}

func SignalPushButton(b *gtk.Builder, buttonID string) error {
	// Инициализация сигнала нажатия кнопки
	obj, err := b.GetObject(buttonID)
	if err != nil {
		return fmt.Errorf("Инициализация нажатия кнопки: %s\n", err)
	}
	button := obj.(*gtk.Button)

	_, err = button.Connect("clicked", func() {
		fmt.Printf("Получена кнопка\n")
	})
	if err != nil {
		return fmt.Errorf("Получение селектора кнопки: %s\n", err)
	}
	return nil
}
