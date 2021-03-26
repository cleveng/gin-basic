package main

import (
	"github.com/maus/basic/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.InitRouter()
	http := &http.Server{
		Addr:           "localhost:8081", // MacOS防火墙会认为启动的非localhost的主机头都是恶意连接（特别是主动启动的）
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	http.ListenAndServe()
}
