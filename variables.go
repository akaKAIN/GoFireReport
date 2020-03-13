package main

import "fmt"

// Количество всех искомых элементов в окне
const (
	VarCount = 10
	EntryCount = 10
	ButtonCount = 10
	LabelCount = 10

)


// Словарь всех полученных элементов страницы, где ключом является ID елемента.
var WindowElements = make(ElemMap, VarCount)

func GenerateID() {
	var  EntryList, ButtonList, LabelList []string
	for i:=1; i<=EntryCount; i++{
		EntryList = append(EntryList, fmt.Sprintf("entry_%d", i))
	}
	for i:=1; i<=ButtonCount; i++{
		EntryList = append(ButtonList, fmt.Sprintf("button_%d", i))
	}
	for i:=1; i<=LabelCount; i++{
		EntryList = append(LabelList, fmt.Sprintf("label_%d", i))
	}

}


