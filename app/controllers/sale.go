package controllers

import (
	"bareskin-api/app/models/request"
	"github.com/revel/revel"
	"bareskin-api/app/models/db"
	"bareskin-api/app/models/response"
	"bareskin-api/app/models/domain"
)

type SaleCtrl struct {
	GorpController
}

func (c SaleCtrl) parseSaleRequest() (request.SaleRequest, string, error) {
	saleRequest := request.SaleRequest{}
	userId := c.Params.Get("id")
	err := c.Params.BindJSON(&saleRequest)
	return saleRequest, userId, err
}

func (c SaleCtrl) Submit() revel.Result {
	if saleRequest, userId, err := c.parseSaleRequest(); err != nil {
		return c.RenderText("Unable to parse SaleRequest from JSON")
	} else {
		sale := db.Sale{}

		sale.UserId = parseIntOrDefault(userId, -1)
		sale.ReferenceNumber = Rand(4) + "-" + Rand(6) + "-" + Rand(4)
		sale.BuyerName = saleRequest.BuyerName
		sale.BuyerNumber = saleRequest.BuyerNumber
		sale.SaleDate, _ = ParseStringToDate(saleRequest.SaleDate)

		if err := c.Txn.Insert(&sale); err != nil {
			return c.RenderText("Error inserting record into database!")
		} else {
			soldProduct := db.SoldProducts{}
			for _, orderUnit := range saleRequest.Purchases {
				soldProduct.SaleId = sale.Id
				soldProduct.ProductId = orderUnit.Product.ProductId
				soldProduct.Quantity = orderUnit.Quantity
				c.Txn.Insert(&soldProduct)
			}
			response := response.SaleResponse{}
			response.StatusCode = 200
			response.Request = saleRequest

			return c.RenderJSON(response)
		}
	}
}

func (c SaleCtrl) History(id int64, sales []db.Sale) revel.Result {

	_, err := c.Txn.Select(&sales,
		`SELECT * FROM bareskindb.Sale WHERE user_id = ?`, id)

	for _, sale := range sales {
		revel.INFO.Printf("%s", sale)
		//get soldProducts table
	}

	if err != nil {
		return c.RenderJSON(domain.JSONError(err.Error()))
	}
	return c.RenderJSON(sales)
}
