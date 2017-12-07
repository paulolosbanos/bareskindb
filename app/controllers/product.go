package controllers

import (
	"bareskin-api/app/models/request"
	"github.com/revel/revel"
	"bareskin-api/app/models/response"
	"bareskin-api/app/models/db"
	"bareskin-api/app/models/domain"
)

type ProductCtrl struct {
	GorpController
}

func (c ProductCtrl) parseProductRequest() (request.ProductRequest, error) {
	productRequest := request.ProductRequest{}
	err := c.Params.BindJSON(&productRequest)
	return productRequest, err
}

func (c ProductCtrl) Add() revel.Result {
	if productRequest, err := c.parseProductRequest(); err != nil {
		return c.RenderText("Unable to parse ProductRequest from JSON")
	} else {
		for _, product := range productRequest.Products {
			c.Txn.Insert(&product)
		}
		response := response.ProductResponse{}
		response.StatusCode = 200
		response.Request = productRequest

		return c.RenderJSON(response)
	}
}

func (c ProductCtrl) List() revel.Result {
	products, err := c.Txn.Select(db.Product{}, `SELECT * FROM Product`)
	if err != nil {
		return c.RenderJSON(domain.JSONError(err.Error()))
	}
	return c.RenderJSON(products)
}
