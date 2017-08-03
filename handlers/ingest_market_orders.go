package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/regner/albiondata-backend/dispatcher"
)

type ingestMarketOrders struct {
	Orders []ingestMarketOrder `json:"orders"`
}

type ingestMarketOrder struct {
	ID               string `json:"Id"`
	ItemID           string `json:"ItemTypeId"`
	LocationID       int    `json:"LocationId"`
	QualityLevel     int    `json:"QualityLevel"`
	EnchantmentLevel int    `json:"EnchantmentLevel"`
	Price            int    `json:"UnitPriceSilver"`
	Amount           int    `json:"Amount"`
	AuctionType      string `json:"AuctionType"`
	Expires          string `json:"Expires"`
}

func IngestMarketOrders(c *gin.Context) {
	var incomingRequest ingestMarketOrders
	c.BindJSON(&incomingRequest)

	for _, v := range incomingRequest.Orders {
		work := dispatcher.Work{
			Topic:   "market-order",
			Message: []byte(json.Marshal(v)),
		}

		dispatcher.WorkQueue <- work
	}
}
