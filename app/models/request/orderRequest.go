package request

import (
	"bareskin-api/app/models/domain"
)

type OrderRequest struct {
	Name           string             `json:"name"`
	Address        string             `json:"address"`
	Mobile         string             `json:"mobile"`
	PaymentChannel string             `json:"payment-channel"`
	PaymentDate    string             `json:"payment-date"`
	OrderUnitList  []domain.OrderUnit `json:"order-unit-list"`
}
