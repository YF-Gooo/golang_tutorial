package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router := gin.Default()
	router.StaticFS("/", http.Dir("D:\\confidenical"))
	// Listen and serve on 0.0.0.0:8080
	router.Run("0.0.0.0:8080")
}
