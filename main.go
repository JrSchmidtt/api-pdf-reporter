package main

import (
	"fmt"
	"pdf-reporter/htmlParser"
)

type Data struct {
	Name string
}

func main(){
	h := htmlParser.New("tmp")

	dataHTML := Data{
		Name : "Junior Schmidt",
	}
	htmlGenerated, err := h.Create("templates/example.html", dataHTML)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(htmlGenerated)
}