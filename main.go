package main

import (
	"github.com/gin-gonic/gin"
	"github.com/harryprasetyo4000/go-restapi-test.git/controllers/productcontroller"
	"github.com/harryprasetyo4000/go-restapi-test.git/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/product/:id", productcontroller.Show)
	r.POST("/api/product", productcontroller.Create)
	r.PUT("/api/product/:id", productcontroller.Update)
	r.DELETE("/api/product", productcontroller.Delete)
	r.POST("/testing/upstream", productcontroller.TestingUp)

	r.Run(":8080")
}
