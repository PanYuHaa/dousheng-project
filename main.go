package main

import (
	"dousheng-demo/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.Init()

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
