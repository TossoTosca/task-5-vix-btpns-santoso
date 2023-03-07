package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Membuat instance dari router
	router := gin.Default()

	// Menambahkan endpoint GET pada path "/"
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, world!",
		})
	})

	// Menjalankan server dengan menggunakan router
	router.Run(":8080")
}
