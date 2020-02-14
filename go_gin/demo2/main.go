package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/", http.Dir("D:\\confidenical"))
	// Listen and serve on 0.0.0.0:8080
	router.Run("0.0.0.0:8080")
}
