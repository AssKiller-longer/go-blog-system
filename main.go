package main

import (
	"log"
	"net/http"
)

func main() {
	//程序入口，一个项目只有一个入口
	//web程序，http协议 ip 端口port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

}
