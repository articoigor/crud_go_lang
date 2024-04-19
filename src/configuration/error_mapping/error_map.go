package error_map

type CurrError struct {
	message string   `json: message`
	err     string   `json: err`
	code    int64    `json: code`
	causes  []Causes `json: causes`
}

type Causes struct {
	field string `json: field`
	cause string `json: cause`
}

func (r *CurrError) Error() string {
	return r.message
}

func NewError(message string, err string, code int64, causes []Causes) *CurrError {
	return &CurrError{
		message: message,
		err:     err,
		code:    code,
		causes:  causes,
	}
}
