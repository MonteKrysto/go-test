package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	db gorm.DB
}

func main() {
	app := App{}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	r.Run()
	
	return nil
}
