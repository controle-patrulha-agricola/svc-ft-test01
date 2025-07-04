// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("main.go svc-FT-test1")

	r.GET("/echo/:text", func(c *gin.Context) {
		input := c.Param("text")
		result := input + "cpa"

		c.String(http.StatusOK, result)
	})

	r.Run(":8888") // Porta 8888
}
