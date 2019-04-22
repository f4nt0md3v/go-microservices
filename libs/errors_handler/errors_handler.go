// Package implement map of ALL client errors.
package errors_handler

type Error interface {
	Error() string
	Code() int
	Details() map[string]interface{}
}

type errorIntance struct {
	CodeMessage int
	Message     string
	DetailsInfo map[string]interface{}
}

func (self errorIntance) Error() string {
	return self.Message
}

func (self errorIntance) Code() int {
	return self.CodeMessage
}

func (self errorIntance) Details() map[string]interface{} {
	return self.DetailsInfo
}

func GetError(code int, details ...map[string]interface{}) Error {
	field, okField := Handler[code]
	if !okField {
		if len(details) > 0 {
			return errorIntance{
				CodeMessage: 10,
				Message:     Handler[10],
				DetailsInfo: details[0],
			}
		} else {
			return errorIntance{
				CodeMessage: 10,
				Message:     Handler[10],
			}
		}
	} else {
		if len(details) > 0 {
			return errorIntance{
				CodeMessage: code,
				Message:     field,
				DetailsInfo: details[0],
			}
		} else {
			return errorIntance{
				CodeMessage: code,
				Message:     field,
			}
		}
	}
}
