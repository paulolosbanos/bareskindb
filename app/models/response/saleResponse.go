package response

import (
	"bareskin-api/app/models/domain"
	"bareskin-api/app/models/request"
)

type SaleResponse struct {
	domain.BaseResponse
	Request request.SaleRequest `json:"request"`
}