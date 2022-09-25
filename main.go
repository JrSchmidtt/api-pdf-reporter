package main

import (
	"fmt"
	"os"
	"pdf-reporter/htmlParser"
	"pdf-reporter/pdfGenerator"
)

type Data struct {
	Name string
}

func main(){
	dataHtml := Data{
		Name : "Lorem Ipsum",
	}

	h := htmlParser.New("tmp")
	p := pdfGenerator.NewWkHtmlToPdf("tmp")

	fileHtml, err := h.Create("templates/example.html", dataHtml)
	defer os.Remove(fileHtml)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File Html: ",fileHtml)

	filePDFName, err := p.Create(fileHtml)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("File Pdf: ",filePDFName)
}