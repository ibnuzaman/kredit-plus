package model

type BaseResponse struct {
	Code    int         `json:"-"`
	Message string      `json:"message" example:"Message!"`
	Data    interface{} `json:"data"`
}

func (res *BaseResponse) Error() string {
	return res.Message
}

func NewErrorMessage(code int, message string, data interface{}) *BaseResponse {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
