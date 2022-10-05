package main

import (
	"github.com/JrSchmidtt/api-gin/server"
)

func main(){
	server := server.NewServer()
	server.Run()
}