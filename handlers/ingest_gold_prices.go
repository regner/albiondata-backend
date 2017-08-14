package handlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/regner/albiondata-backend/dispatcher"
)

type ingestGoldPrices struct {
	Prices     []int `json:"GoldPrices"`
	TimeStamps []int `json:"TimeStamps"`
}

func IngestGoldPrices(c *gin.Context) {
	var incomingRequest ingestGoldPrices
	c.BindJSON(&incomingRequest)

	data, _ := json.Marshal(incomingRequest)

	work := dispatcher.Work{
		Topic:   "gold-prices",
		Message: []byte(data),
	}

	dispatcher.WorkQueue <- work
}
