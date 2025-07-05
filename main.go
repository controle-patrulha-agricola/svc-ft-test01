// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	fmt.Println("================= main.go svc-FT-test1")

	r.GET("/echo/:text", func(c *gin.Context) {

		input := c.Param("text")
		result := "\n Teste do cpa: vc digitou " + input + "\n"
		c.String(http.StatusOK, result)
		fmt.Printf("================= main.go handler /echo: %s\n", input)
	})

	r.Run(":9999") // Porta 9999
}
