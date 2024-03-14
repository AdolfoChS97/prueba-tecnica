package main

import (
	CSV "api/packages/csv"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.GET("/csv", func(c *gin.Context) {
		csvData := CSV.ReadCsvFile("../data/customers-100000.csv")
		c.JSON(http.StatusOK, gin.H{
			"message": CSV.PaginatesCsvFile(csvData, 2, 10, len(csvData)),
		})
	})

	router.Run()
}
