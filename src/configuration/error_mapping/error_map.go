package error_map

type CurrError struct {
	Message string  `json:"message"`
	Err     string  `json:"err"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes"`
}

type Cause struct {
	Field string `json:"field"`
	Cause string `json:"cause"`
}

func (r *CurrError) Error() string {
	return r.Message
}

func NewError(message string, err string, code int, causes []Cause) *CurrError {
	return &CurrError{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewErrorBadRequest(message string) *CurrError {
	return &CurrError{
		Message: message,
		Err:     "bad_request",
		Code:    404,
	}
}

func NewErrorValidationBadRequest(message string, causes []Cause) *CurrError {
	return &CurrError{
		Message: message,
		Err:     "bad_request",
		Code:    404,
		Causes:  causes,
	}
}
