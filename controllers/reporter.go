package controllers

import (
	"fmt"
	"os"

	"github.com/JrSchmidtt/api-gin/service"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Name string
}

func CreatePDF(c *gin.Context) {

	dataHtml := Data{
		Name: "Lorem Ipsum",
	}

	if err := service.EnsureDir("./tmp"); err != nil {
		c.JSON(500, gin.H{
			"error": "Directory creation failed with error:" + err.Error(),
		})
		return
	}

	h := service.NewHtml("./tmp")
	p := service.NewWkHtmlToPdf("./tmp")

	fileHtml, err := h.CreateHtml("templates/example.html", dataHtml)
	defer os.Remove(fileHtml)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "" + err.Error(),
		})
		return
	}
	fmt.Println("File Html: ", fileHtml)

	filePDFName, err := p.CreatePDF(fileHtml)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "" + err.Error(),
		})
		return
	}
	fmt.Println("File Pdf: ", filePDFName)
	//defer os.Remove(fileHtml)

	c.JSON(200, gin.H{
		"File Html": fileHtml,
		"File Pdf":  filePDFName,
	})
}
