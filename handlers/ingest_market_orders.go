package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/regner/albiondata-backend/dispatcher"
)

type ingestMarketOrders struct {
	Orders     []ingestMarketOrder `json:"Orders"`
	LocationID int                 `json:"LocationId"`
}

type ingestMarketOrder struct {
	ID               int    `json:"Id"`
	ItemID           string `json:"ItemTypeId"`
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
		data, _ := json.Marshal(v)

		work := dispatcher.Work{
			Topic:   "market-order",
			Message: []byte(data),
		}

		dispatcher.WorkQueue <- work
	}
}
