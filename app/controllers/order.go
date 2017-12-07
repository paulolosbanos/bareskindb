package controllers

import (
	"github.com/revel/revel"
	"bareskin-api/app/models/request"
	"bareskin-api/app/models/db"
	"time"
	"bareskin-api/app/models/response"
)

type OrderCtrl struct {
	GorpController
}

func (c OrderCtrl) parseOrderRequest() (request.OrderRequest, string, error) {
	orderRequest := request.OrderRequest{}
	userId := c.Params.Get("id")
	err := c.Params.BindJSON(&orderRequest)
	return orderRequest, userId, err
}

func (c OrderCtrl) Submit() revel.Result {
	if orderRequest, userId, err := c.parseOrderRequest(); err != nil {
		return c.RenderText("Unable to parse OrderRequest from JSON")
	} else {
		//add Validation later

		order := db.Order{}

		//id, _ := strconv.ParseInt(userId, 10, 64)
		order.UserId = parseIntOrDefault(userId, -1)
		order.Name = orderRequest.Name
		order.Address = orderRequest.Address
		order.Mobile = orderRequest.Mobile
		order.PaymentChannel = orderRequest.PaymentChannel
		order.PaymentDate, _ = ParseStringToDate(orderRequest.PaymentDate)
		order.ReferenceNumber = Rand(16)
		order.OrderDate = time.Now()
		order.OrderStatus = "Awaiting for confirmation"

		if err := c.Txn.Insert(&order); err != nil {
			return c.RenderText(
				"Error inserting record into database!")
		} else {
			orderedProducts := db.OrderedProducts{}
			for _, val := range orderRequest.OrderUnitList {
				orderedProducts.OrderId = order.Id
				orderedProducts.ProductId = val.Product.ProductId
				orderedProducts.Quantity = val.Quantity
				c.Txn.Insert(&orderedProducts)
			}

			response := response.OrderResponse{}
			response.StatusCode = 200
			response.ReferenceNumber = order.ReferenceNumber
			response.OrderDate = order.OrderDate.String()
			response.OrderStatus = order.OrderStatus
			response.Request = orderRequest

			return c.RenderJSON(response)
		}
	}
}
