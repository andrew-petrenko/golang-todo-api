package base_response

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func NewResponse(data interface{}, success bool) *Response {
	return &Response{
		Success: success,
		Data:    data,
	}
}

func NewResponseErrorMessage(message string) *Response {
	return &Response{
		Success: false,
		Data:    map[string]string{"message": message},
	}
}
