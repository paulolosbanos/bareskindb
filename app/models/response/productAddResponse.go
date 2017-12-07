package response

import (
	"bareskin-api/app/models/domain"
	"bareskin-api/app/models/request"
)

type ProductResponse struct {
	domain.BaseResponse
	Request request.ProductRequest `json:"request"`
}
