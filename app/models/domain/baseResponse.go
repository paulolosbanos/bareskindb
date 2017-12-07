package domain

type BaseResponse struct {
	StatusCode    int `json:"status-code"`
	StatusMessage string `json:"status-message"`
}
