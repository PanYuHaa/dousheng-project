package main

import (
	"dousheng-demo/repository"
	"dousheng-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.Init()
	service.Init()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
