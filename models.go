package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
)

// Словарь где ключи - ID элементов, а значения - сами элементы окна
type ElemMap map[string]Element

// Метод добавления элемента в словарь
func (em *ElemMap) Add(element Element){
	(*em)[element.Id] = element
}

type Element struct {
	Name       string
	Id         string
	ElemType   interface{}
	Signal     string
	Builder    *gtk.Builder
}

func (e *Element) GetEntry() (*gtk.Entry, error) {
	//Метод получения поля ввода по ID структуры
	obj, err := e.Builder.GetObject(e.Id)
	if err != nil {
		err = fmt.Errorf("Ошибка получения поля ввода %q. Error: %s\n", e.Id, err)
		return nil, err
	}
	return obj.(*gtk.Entry), nil
}

func (e *Element) GetButton() (*gtk.Button, error) {
	//Метод получения кнопки по ID структуры
	obj, err := e.Builder.GetObject(e.Id)
	if err != nil {
		err = fmt.Errorf("Ошибка получения кнопки %q. Error: %s\n", e.Id, err)
		return nil, err
	}
	return obj.(*gtk.Button), nil
}

func (e *Element) GetLabel() (*gtk.Label, error) {
	//Метод получения текстового поля по ID структуры
	obj, err := e.Builder.GetObject(e.Id)
	if err != nil {
		err = fmt.Errorf("Ошибка получения текстового поля %q. Error: %s\n", e.Id, err)
		return nil, err
	}
	return obj.(*gtk.Label), nil
}

func (e *Element) GetTextArea() (*gtk.TextView, error){
	//Метод получения текстового поля по ID структуры
	obj, err := e.Builder.GetObject(e.Id)
	if err != nil {
		err = fmt.Errorf("Ошибка получения text %q. Error: %s\n", e.Id, err)
		return nil, err
	}
	return obj.(*gtk.TextView), nil

}

func (e *Element) GetTextBuffer() (*gtk.TextBuffer, error){
	//Метод получения текстового буфера по ID структуры
	obj, err := e.Builder.GetObject(e.Id)
	if err != nil {
		err = fmt.Errorf("Ошибка получения text %q. Error: %s\n", e.Id, err)
		return nil, err
	}
	return obj.(*gtk.TextBuffer), nil
}

