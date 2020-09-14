package main

import (
	"github.com/gin-gonic/gin"
	"github.com/serkancetintas/go-book/controllers"
	"github.com/serkancetintas/go-book/models"
)

func main() {
	r := gin.Default()

	db := models.SetupModels()

	// Provide db variable to controllers
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)  // update by id
	r.DELETE("/books/:id", controllers.DeleteBook) // delete by id

	r.Run()
}
