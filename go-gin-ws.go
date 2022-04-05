package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id       int64   `json:"id"`
	Name     string  `json:"name"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
}

var products = []Product{
	{1, "Notebook", 15, 2500.00},
	{2, "TV", 22, 2200.00},
	{3, "Smartphone", 125, 1500.00},
}

func main() {
	router := gin.Default()
	router.GET("/products", getProducts)
	router.GET("/products/:id", getProduct)
	router.POST("/products", createProduct)
	router.Run("localhost:8080")
}

func getProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, products)
}

func getProduct(context *gin.Context) {
	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	for _, value := range products {
		if value.Id == id {
			context.IndentedJSON(http.StatusOK, value)
		}
	}
}

func createProduct(context *gin.Context) {
	var product Product
	err := context.BindJSON(&product)
	if err != nil {
		fmt.Println("An error occurred while trying to create a new Product")
	}
	products = append(products, product)
	context.IndentedJSON(http.StatusCreated, products)
}
