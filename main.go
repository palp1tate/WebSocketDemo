package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	registerRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func registerRoutes(r *gin.Engine) {
	r.GET("/chat", ChatHandler)
}
