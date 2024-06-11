package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patepath/gglib/usercontroller"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("start...")

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var v1 = r.Group("/v1")
	{
		v1.GET("/user/:id", usercontroller.FindOne)
		v1.GET("/user", usercontroller.FindAll)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
