package main

import (
	"fmt"
	"github.com/gotk3/gotk3/gtk"
	"io/ioutil"
	"os"
)

// Количество всех искомых элементов в окне
const (
	VarCount    = 4
	EntryCount  = 1
	ButtonCount = 2
	LabelCount  = 2
)

func GenerateID() {
	var EntryList, ButtonList, LabelList []string
	for i := 1; i <= EntryCount; i++ {
		EntryList = append(EntryList, fmt.Sprintf("entry_%d", i))
	}
	for i := 1; i <= ButtonCount; i++ {
		EntryList = append(ButtonList, fmt.Sprintf("button_%d", i))
	}
	for i := 1; i <= LabelCount; i++ {
		EntryList = append(LabelList, fmt.Sprintf("label_%d", i))
	}

}

func InitGlobalList(builder *gtk.Builder) {
	// Словарь всех полученных элементов страницы, где ключом является ID елемента.
	//var WindowElements = make(ElemMap, VarCount)
	var (
		createButton = Element{
			Name:     "create",
			Id:       "create",
			ElemType: "button",
			Signal:   "clicked",
			Builder:  builder,
		}

		resultLabel = Element{
			Name:     "label",
			Id:       "label",
			ElemType: "label",
			Builder:  builder,
		}

		inputField = Element{
			Name:     "input",
			Id:       "input",
			ElemType: "input",
			Builder:  builder,
		}
		textArea = Element{
			Name:     "text_area",
			Id:       "text_area",
			ElemType: "area",
			Builder:  builder,
		}
		textBuffer = Element{
			Name:     "textbuffer1",
			Id:       "textbuffer1",
			ElemType: "textBuffer",
			Builder:  builder,
		}

		FList      []os.FileInfo
		resultText string

	)

	btnCreate, err := createButton.GetButton()
	MyLogger.CheckErr("", err)

	Label, err := resultLabel.GetLabel()
	MyLogger.CheckErr("", err)

	input, err := inputField.GetEntry()
	MyLogger.CheckErr("", err)

	dirList, err := textArea.GetTextArea()
	MyLogger.CheckErr("", err)

	buf, err := textBuffer.GetTextBuffer()
	MyLogger.CheckErr("", err)

	//w, err := builder.GetObject("viewport")
	//MyLogger.CheckErr("", err)

	_, err = btnCreate.Connect(createButton.Signal, func() {
		inputPath, err := input.GetText()
		MyLogger.CheckErr("Get text from input field", err)

		fmt.Println("Get text: ", inputPath)

		if err == nil {
			Label.SetText(inputPath)
			input.SetText("")
			FList = GetDirList(inputPath)
			resultText = ""
			for _, f := range FList {
				resultText = resultText + f.Name() + "\n"
			}
		}
		buf.SetText(resultText)
		dirList.SetBuffer(buf)
		fmt.Println("Done")
	})
}

func GetDirList(path string) []os.FileInfo {
	result, err := ioutil.ReadDir(path)
	MyLogger.CheckErr("Reading directory: "+path, err)
	return result
}
