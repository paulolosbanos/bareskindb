package request

import "bareskin-api/app/models/domain"

type SaleRequest struct {
	BuyerName   string             `json:"buyer-name"`
	BuyerNumber string             `json:"buyer-number"`
	SaleDate    string             `json:"sale-date"`
	Purchases   []domain.OrderUnit `json:"purchases"`
}
