package service

import (
	"os"
	"text/template"

	"github.com/JrSchmidtt/api-gin/models"
	"github.com/google/uuid"
)

type htmlStruct struct {
	rootPath string
}

func NewHtml(rootPath string) models.HTMLParserInterface {
	return &htmlStruct{rootPath: rootPath}
}

func (a *htmlStruct) CreateHtml(templateName string, data interface{}) (string, error) {
	templateGenerator, err := template.ParseFiles(templateName)
	if err != nil {
		return "", err
	}

	fileName := a.rootPath + "/" + uuid.New().String() + ".html"

	fileWriter, err := os.Create(fileName)
	if err != nil {
		return "", err
	}

	if err := templateGenerator.Execute(fileWriter, data); err != nil{
		return "", nil
	}


	return fileName, nil
}