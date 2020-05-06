package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type ResponseHelper struct {
}

//response format
type responseFormat struct {
	C       echo.Context
	Code    int
	Status  string
	Message string
	Data    interface{}
}

// SetResponse ...
func (r *ResponseHelper) SetResponse(c echo.Context, code int, status string, message string, data interface{}) responseFormat {
	return responseFormat{c, code, status, message, data}
}

// SendResponse ...
func (r *ResponseHelper) SendResponse(res responseFormat) error {
	if len(res.Message) == 0 {
		res.Message = http.StatusText(res.Code)
	}

	if res.Data != nil {
		return res.C.JSON(res.Code, map[string]interface{}{
			"code":    res.Code,
			"status":  res.Status,
			"message": res.Message,
			"data":    res.Data,
		})
	} else {
		return res.C.JSON(res.Code, map[string]interface{}{
			"code":    res.Code,
			"status":  res.Status,
			"message": res.Message,
		})
	}
}

// EmptyJSONMap : set empty data.
func (r *ResponseHelper) EmptyJSONMap() map[string]interface{} {
	return make(map[string]interface{})
}

// SendSuccess : Send success response to consumers.
func (r *ResponseHelper) SendSuccess(c echo.Context, message string, data interface{}) error {
	res := r.SetResponse(c, http.StatusOK, "success", message, data)
	return r.SendResponse(res)
}

// SendBadRequest : Send bad request response to consumers.
func (r *ResponseHelper) SendBadRequest(c echo.Context, message string, data interface{}) error {
	res := r.SetResponse(c, http.StatusBadRequest, "error", message, data)
	return r.SendResponse(res)
}

// SendError : Send error request response to consumers.
func (r *ResponseHelper) SendError(c echo.Context, message string, data interface{}) error {
	res := r.SetResponse(c, http.StatusInternalServerError, "error", message, data)
	return r.SendResponse(res)
}

// SendUnauthorized : Send error request response to consumers.
func (r *ResponseHelper) SendUnauthorized(c echo.Context, message string, data interface{}) error {
	res := r.SetResponse(c, http.StatusUnauthorized, "error", message, data)
	return r.SendResponse(res)
}

// SendValidationError : Send validation error request response to consumers.
func (r *ResponseHelper) SendValidationError(c echo.Context, validationErrors validator.ValidationErrors) error {
	errorResponse := []string{}
	for _, err := range validationErrors {
		errorResponse = append(errorResponse, strings.Trim(fmt.Sprint(err), "[]")+".")
	}
	res := r.SetResponse(c, http.StatusBadRequest, "error", strings.Trim(fmt.Sprint(errorResponse), "[]"), r.EmptyJSONMap())
	return r.SendResponse(res)
}
