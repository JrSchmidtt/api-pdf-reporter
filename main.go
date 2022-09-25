package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"pdf-reporter/htmlParser"
	"pdf-reporter/pdfGenerator"
)

type Data struct {
	Name string
}

type Response struct {
	Id int
	FileName string
}

func main(){
	http.HandleFunc("/pdf", getPDF)
	fmt.Println("Api Running..")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func getPDF(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content=Type", "application/json")

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
	
	json.NewEncoder(w).Encode([]Response{{
		Id: 1,
		FileName: filePDFName,
	}})
}