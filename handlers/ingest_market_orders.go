package handlers

import (
"github.com/gin-gonic/gin"
"github.com/regner/albiondata-backend/dispatcher"
)

type ingestMarketOrders struct {
	Orders []string `json:"orders"`
}

type ingestMarketOrder struct {
	ID string `jston:"Id"`
	order_id=order_json['Id'],
item_id=item_id,
location_id=args['LocationId'],
quality=order_json['QualityLevel'],
enchantment=order_json['EnchantmentLevel'],
price=price,
amount=order_json['Amount'],
expire=dateutil.parser.parse(order_json['Expires']),
is_buy_order=is_buy_order,
}

func IngestMarketOrders(c *gin.Context) {
	var incomingRequest ingestMarketOrders
	c.BindJSON(&incomingRequest)

	for _, v := range incomingRequest.Orders {
		work := dispatcher.Work{
			Topic:   "market-order",
			Message: []byte(v),
		}

		dispatcher.WorkQueue <- work
	}
}
