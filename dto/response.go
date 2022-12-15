package dto

type Response struct {
	Code	int 	`json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(statusCode int, message string, data any) Response {
	return Response{
		Code: statusCode,
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(statusCode int, message string, err error) Response {
	return Response{
		Code: statusCode,
		Message: message,
		Data:    err,
	}
}

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
