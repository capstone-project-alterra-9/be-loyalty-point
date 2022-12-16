package dto

type ResponseSuccess struct {
	Code	 	int		`json:"code"`
	Message 	string 		`json:"message"`
	Data    	any    		`json:"data"`
}

type ResponseError struct {
	Code	 	int		`json:"code"`
	Message  	string 		`json:"message"`
	ErrorId  	string    	`json:"errors"`
	Data     	any    		`json:"data"`
}

type EmptyObj struct{}

func BuildResponse(message string, data any) ResponseSuccess {
	return ResponseSuccess{
		Code:	200,
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err error) ResponseError {
	// errorCode := err.Error()
	return ResponseError{
		Code:	200,
		Message: message,
		ErrorId:  "error id",
		Data:    err,
	}
}

type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
