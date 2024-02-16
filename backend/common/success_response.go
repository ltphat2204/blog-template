package common

type successResponse struct {
	Successs   bool        `json:"success"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

func NewSuccessResponse(data interface{}, pagination Pagination) *successResponse {
	return &successResponse{
		Successs:   true,
		Data:       data,
		Pagination: &pagination,
	}
}

func NewSimpleSuccessResponse(data interface{}) *successResponse {
	return &successResponse{
		Successs: true,
		Data:     data,
	}
}
