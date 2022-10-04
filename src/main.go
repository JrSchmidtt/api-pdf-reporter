package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pdf-reporter/htmlParser"
	"pdf-reporter/pdfGenerator"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Data struct {
	Name string
}

type Response struct {
	Id int
	FileName string
}

var (
	client *s3.S3
)

func init(){
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET_KEY"),""),
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatal(err)
	}
	client = s3.New(sess)
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

	h := htmlParser.New("./")
	p := pdfGenerator.NewWkHtmlToPdf("./")

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

	fmt.Println("Upload start!")
	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String("CHANGE-BUCKET-NAME-HERE"),
		Key: aws.String(filePDFName),
		Body: bytes.NewReader(file_bytes),
	})
	if err != nil{
		fmt.Printf("Upload error: %v",err)
	}
	fmt.Println("Upload end!")

	w.WriteHeader(http.StatusOK)
	w.Write(file_bytes)
    defer fileD.Close()
	defer os.Remove(filePDFName)
}