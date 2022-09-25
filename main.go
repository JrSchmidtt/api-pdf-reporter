package main

import (
	"fmt"
	"io/ioutil"
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
	w.Header().Add("content-type", "application/pdf")

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

	fileD, err := os.Open(filePDFName)
    if err != nil {
        log.Panic(err)
    }
	
    file_bytes, err := ioutil.ReadAll(fileD)
    if err != nil {
        log.Panic(err)
    }
	w.WriteHeader(http.StatusOK)
	w.Write(file_bytes)
    defer fileD.Close()
	defer os.Remove(filePDFName)
}