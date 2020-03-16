package main

import (
	"github.com/gotk3/gotk3/gtk"
	"path/filepath"
)

func main() {
	templateName := "create_file_form.glade"
	// Задаем путь до шаблона окна
	template := filepath.Join("source", "templates", templateName)

	// Инициализируем GTK.
	gtk.Init(nil)

	// Создаем билдер
	builder, err := gtk.BuilderNew()
	MyLogger.CheckErr("Инициализация билдера", err)

	// Загружаем билдер в окно из файла Glade
	err = builder.AddFromFile(template)
	MyLogger.CheckErr("Открытие формы", err)

	InitGlobalList(builder)

	// Получаем главное окно
	obj, err := builder.GetObject("main_window")
	MyLogger.CheckErr("Получение селектора главного окна", err)

	// Преобразуем обект в окно типа gtk.Window
	// и соединяем с сигналом "destroy", чтобы приложение закрывалось при
	// закрытии окна.
	win := obj.(*gtk.Window)
	win.SetTitle("InfoDev: Excel Aggregate ©")

	go func() {
		err = SignalCloseWindow(win)
		MyLogger.CheckErr("", err)
	}()

	win.SetDefaultSize(350, 280)
	// Отображение всех виджетов в окне
	win.ShowAll()

	// Выполняем главнй цикл для орисовки окна.
	// Он завершится когда выполнится gtk.MainQuit
	gtk.Main()

}
