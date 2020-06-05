package main

import (
	"github.com/izabela-k/lesfinder/db"
	"github.com/izabela-k/lesfinder/users"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(db.Postgres())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/sign-up/", users.SignUpHandler)
	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
