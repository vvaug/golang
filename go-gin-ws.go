package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Name     string  `json:"name"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
}

var products = []Product{
	{"Notebook", 15, 2500.00},
	{"TV", 22, 2200.00},
	{"Smartphone", 125, 1500.00},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.Run("localhost:8080")
}

func getProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, products)
}
