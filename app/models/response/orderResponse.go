package response

import (
	"bareskin-api/app/models/request"
	"bareskin-api/app/models/domain"
)

type OrderResponse struct {
	domain.BaseResponse
	ReferenceNumber string               `json:"reference-number"`
	OrderDate       string               `json:"order-date"`
	OrderStatus     string               `json:"order-status"`
	Request         request.OrderRequest `json:"request"`
}
