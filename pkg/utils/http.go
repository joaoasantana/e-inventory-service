package utils

type (
	StatusResponse struct {
		Code    int    `json:"code"    extensions:"x-order:0"`
		Message string `json:"message" extensions:"x-order:1"`
	}

	SuccessResponse struct {
		Status StatusResponse `json:"status" extensions:"x-order:0"`
		Data   interface{}    `json:"data"   extensions:"x-order:1"`
	}

	ErrorResponse struct {
		Status StatusResponse `json:"status" extensions:"x-order:0"`
		Error  string         `json:"error"  extensions:"x-order:1"`
	}
)
