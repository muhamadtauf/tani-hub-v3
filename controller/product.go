package controller

import (
	"net/http"
	"strconv"
	"tani-hub-v3/database"
	"tani-hub-v3/repository"
	"tani-hub-v3/structs"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	var result gin.H

	products, err := repository.GetAllProduct(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": products,
		}
	}
	c.JSON(http.StatusOK, result)
}

func GetProductById(c *gin.Context) {
	var product structs.Product
	id, _ := strconv.Atoi(c.Param("id"))

	product.Id = int64(id)

	var result gin.H

	products, err := repository.GetProductById(database.DbConnection, product)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": products,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertProduct(c *gin.Context) {
	var product structs.Product

	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}

	err = repository.InsertProduct(database.DbConnection, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Product",
	})
}

func UpdateProduct(c *gin.Context) {
	var product structs.Product
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}

	product.Id = int64(id)
	product.UpdatedAt = time.Now()

	err = repository.UpdateProduct(database.DbConnection, product)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Update Product",
	})

}

func DeleteProduct(c *gin.Context) {
	var product structs.Product
	id, err := strconv.Atoi(c.Param("id"))

	product.Id = int64(id)

	err = repository.DeleteProduct(database.DbConnection, product)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Product",
	})
}
