package response

import (
	"github.com/go-playground/validator/v10"
)

type ResponseError struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Success bool        `json:"success"`
}

type WebError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *WebError) Error() string {
	return e.Message
}

func ResponseIfError(err error, code int) {
	if err != nil {
		// Handle error from validators
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errList []string
			for _, errData := range validationErrors {
				err := ("Form validation for " + errData.Field() + " failed on " + errData.Tag() + " " + errData.Param())
				errList = append(errList, err)
			}

			panic(ResponseError{
				Code:    code,
				Message: errList,
				Success: false,
			})
		}

		// Handle error from ResponseError struct
		panic(ResponseError{
			Code:    code,
			Message: err.Error(),
			Success: false,
		})
	}
}

func ValidatorParsers(err error) interface{} {
	var errList []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, errData := range validationErrors {
			err := ("form validation for " + errData.Field() + " failed on " + errData.Tag() + " " + errData.Param())
			errList = append(errList, err)
		}
		return errList
	}
	return err.Error()
}
