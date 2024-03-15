package main

import (
	CSV "api/packages/csv"
	Response "api/packages/response"
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
		csv := CSV.ReadCsvFile("../data/customers-100000.csv")
		records := Response.Paginate(csv.Records, 1, 10, len(csv.Records))
		headers := csv.Headers
		data := Response.Map(headers, records)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    data,
			"total":   len(data) + 1,
			"code":    0,
		})
	})

	router.Run()
}
