package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/viniciusarambul/transaction/src/infra"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "estamos on"})
	})

	_, err := infra.SetupDB()
	if err != nil {
		fmt.Println(err)
		panic("errou")
	}

	engine.Run()
}
