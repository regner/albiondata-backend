package handlers

import (
"github.com/gin-gonic/gin"
"github.com/regner/albiondata-backend/dispatcher"
)

type ingestPostRequest struct {
	Items []string `json:"items"`
}

func IngestMarketOrders(c *gin.Context) {
	var incomingRequest ingestPostRequest
	c.BindJSON(&incomingRequest)

	for _, v := range incomingRequest.Items {
		work := dispatcher.Work{
			Topic:   "market-order",
			Message: []byte(v),
		}

		dispatcher.WorkQueue <- work
	}
}
