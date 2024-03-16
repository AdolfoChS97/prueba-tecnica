package main

import (
	CSV "api/packages/csv"
	Response "api/packages/response"
	"fmt"
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

		var queryParams struct {
			FirstName        string
			LastName         string
			City             string
			Email            string
			SubscriptionDate string
		}

		if err := c.ShouldBindQuery(&queryParams); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "error",
				"error":   err.Error(),
				"code":    1,
			})
			return
		}

		queryParamsData := c.Request.URL.Query()
		fmt.Println(queryParamsData)

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
