package mapper

type ErrorResponse struct {
	Error string `json:"error"`
}

func ToErrorResponse(err string) ErrorResponse {
	return ErrorResponse{
		Error: err,
	}
}

type SimpleResponse struct {
	Message string `json:"message"`
}

func ToSimpleMessageResponse(msg string) SimpleResponse {
	return SimpleResponse{
		Message: msg,
	}
}
