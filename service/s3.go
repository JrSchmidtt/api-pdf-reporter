package service

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


var (
	client *s3.S3
)

func initS3(){
	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY"),
			os.Getenv("AWS_SECRET_KEY"),""),
		Region: aws.String(os.Getenv("AWS_BUCKET_NAME")),
	})
	if err != nil {
		log.Fatal(err)
	}
	client = s3.New(sess)
}

func UploadFileToS3(filePDFName string)(error, string){
	fileD, err := os.Open(filePDFName)
    if err != nil {
        log.Panic(err)
		return err, ""
    }

	file_bytes, err := ioutil.ReadAll(fileD)
    if err != nil {
        log.Panic(err)
		return err, ""
    }

	s3Data, err := client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key: aws.String(filePDFName),
		Body: bytes.NewReader(file_bytes),
	})
	if err != nil{
		fmt.Printf("Upload error: %v",err)
		return err, "Upload error"
	}

	fmt.Printf("S3 DATA: %v", s3Data)

	return nil, "Upload end!"
}