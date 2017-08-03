package main

import (
	"github.com/gin-gonic/gin"
	"github.com/regner/albiondata-backend/dispatcher"
	"github.com/regner/albiondata-backend/handlers"
)

func main() {
	go dispatcher.StartDispatcher()

	r := gin.Default()

	r.POST("/api/v1/ingest/marketorders/", handlers.IngestMarketOrders)

	r.Run(":8080")
}
